const {Kafka} = require('kafkajs');
const sunbirdRegistryService = require('./src/services/sunbird.service');
const config = require('./src/configs/config');

console.log("registry async handler");
const kafka = new Kafka({
    clientId: 'registry-async-handler',
    brokers: config.KAFKA_BOOTSTRAP_SERVER.split(",")
});
const consumer = kafka.consumer({ groupId: 'post_create_entity', sessionTimeout: config.KAFKA_CONSUMER_SESSION_TIMEOUT });

(async function (){
    await consumer.connect();
    await consumer.subscribe({topic: config.POST_CREATE_ENTITY_TOPIC, fromBeginning: true});
    
    console.log("Stored Entity type: ", config.STORED_ENTITY_TYPE);
    console.log("Transaction Entity type: ", config.TRANSACTION_ENTITY_TYPE);

    await consumer.run({
      eachMessage: async ({message}) => {
        const postCreateEntityMessage = JSON.parse(message.value.toString());
        if(postCreateEntityMessage.entityType == config.TRANSACTION_ENTITY_TYPE){
          const addTransactionRequest = sunbirdRegistryService.createTransactionRequest(postCreateEntityMessage);
          console.log({addTransactionRequest: addTransactionRequest});
          try{
            const transactionResponse = await sunbirdRegistryService.addTransaction(addTransactionRequest,config.STORED_ENTITY_TYPE);
            if(transactionResponse?.status == 200){
              console.log("Successfully added Transaction to: ",config.STORED_ENTITY_TYPE);
            }else{
              console.log("Failed to add Transaction", transactionResponse);
            }
          }catch(err){
            console.error("Error in adding transaction",err);
          }
        }
      }
    });
})();

const uuid = require('uuid');

const sunbirdRegistryService = require('../services/sunbird.service')
const certifyConstants = require('../configs/constants');
const {validationResult} = require('express-validator');
const validationService = require('../services/validation.service');
const {truncateShard} = require("../utils/certification.utils");

const REVOKED = "REVOKED";
const SUSPENDED = "SUSPENDED";
const VALID = "VALID";
const INVALID = "INVALID";

async function createCertificate(req, res, kafkaProducer) {
    try {
        validationService.validateCertificateInput(req, "create");
        await kafkaProducer.connect();
        const transactionId = uuid.v4();
        kafkaProducer.send({
            topic: certifyConstants.VC_CERTIFY_TOPIC,
            messages: [
                { key: null, value: JSON.stringify({ body: req.body, transactionId: transactionId, entityType: req.params.entityType, token: req.header("Authorization") }) }
            ]
        });
        res.status(200).json({
            transactionId
        });
    } catch (err) {
        console.error(err);
        res.status(err?.response?.status || 500).json({
            message: err?.response?.data
        });
    }
}

async function getCertificate(req, res) {
    try {
        const entityName = req.params.entityName;
        const certificateId = req.params.certificateId;
        const filters = {
            "filters": {
                "certificateId": {
                    "eq": certificateId
                }
            },
            "limit": 1,
            "offset": 0
        }
        let certificateResponse = await sunbirdRegistryService.searchCertificate(entityName, filters, req.header("Authorization"))
        let certificateOsId = truncateShard(certificateResponse[0]?.osid);
        const {data} = await sunbirdRegistryService.getCertificate(entityName, certificateOsId, req.headers);
        if (req.headers.accept === certifyConstants.SVG_ACCEPT_HEADER) {
            res.type(certifyConstants.IMAGE_RESPONSE_TYPE);
        }
        data.pipe(res);
    } catch (err) {
        console.error(err);
        res.status(err?.response?.status || 500).json({
            message: err
        });
    }
}

async function updateCertificate(req, res, kafkaProducer) {
    try {
        validationService.validateCertificateInput(req, "update");
        await kafkaProducer.connect();
        const transactionId = uuid.v4();
        kafkaProducer.send({
            topic: certifyConstants.VC_CERTIFY_TOPIC,
            messages: [
                { key: null, value: JSON.stringify({ body: req.body, transactionId: transactionId, entityType: req.params.entityType, token: req.header("Authorization") }) }
            ]
        });
        res.status(200).json({
            transactionId
        });
    } catch (err) {
        console.error(err);
        res.status(err?.response?.status || 500).json({
            message: err?.response?.data
        });
    }
}

async function revokeCertificate(req, res) {
    const errors = validationResult(req);
    console.log(errors);
    if (!errors.isEmpty()) {
        res.status(400).json(errors);
        return;
    }
    const token = req.header("Authorization");
    getEntity(req.body.certificateId,"certificateId",req.body.entityName,token)
    .then(async(result) => {
        if(result.length >= 1) {
            let body = getRevokeBody(req);
            const certificateRevokeResponse = await sunbirdRegistryService.revokeCertificate(body, token);
            res.status(200).json({
                message: "Certificate Revoked",
                certificateRevokeResponse: certificateRevokeResponse
            });
        }
        else {
            console.log('RESULT : ',result);
            res.status(400).json({
                message: `Entry for ${req.body.entityName} not found`
            })
        }
    }).catch(err => {
        console.log('ERROR : ',err?.response?.status || '');
        res.status(err?.response?.status || 500).json({
            message: err?.response?.data
        })
    })
}
function getRevokeBody(req) {
    let body = {
        previousCertificateId: req.body.certificateId,
        schema: req.body.entityName,
        startDate: new Date(),
    }
    if (req.body.endDate) {
        body = { ...body, endDate: req.body.endDate }
    }
    return body;

}


async function verifyCertificate (req,res){
    const certificate = req.body;
    const certificateEntityType = certificate.evidence[0].type[0];
    const revokeEntityType = "RevokedVC";
    const token = req.header("Authorization");
    let certificateId= certificate.credentialSubject.id;
    let certificateStatus = "";
    let msg = "";
    console.log({certificateId: certificateId});
    let body = {
        signedCredentials : certificate,
    }
    try{
        const verifyResp = await sunbirdRegistryService.verifyCertificate(body)
        if(verifyResp.verified){
            const certificateResponse = await getEntity(certificateId,"certificateId",certificateEntityType,token);
            if(certificateResponse.length  >= 1){
                const revokeResp = await getEntity(certificateId,"previousCertificateId",revokeEntityType,token);
                const revokeEntityResp = revokeResp.filter(resp =>  resp.schema === certificateEntityType);
                console.log("revokeEntityResp:", revokeEntityResp);
                [certificateStatus,msg] = revokeStatus(revokeEntityResp);
                
            }else{
                certificateStatus = INVALID;
                msg = `Certificate is not available in ${certificateEntityType}`
            }
            res.status(200).json({
                message: "Certificate verified",
                status: {
                    certificateStatus: certificateStatus,
                    msg: msg
                },
                response: {verifyResp}
            });
        }else{
            res.status(406).json({
                message: "verification failed",
                status: {
                    certificateStatus: INVALID,
                    msg: "Failed to verify certificate"
                },
                response: {verifyResp}
            });
        }
    } catch(err) {
        console.log('ERROR : ',err);
        res.status(500).json({
            message: err?.response?.data
        })
    }
}

function revokeStatus(revokeEntityResp){
    let certificateStatus = "";
    let msg = "";
    if(revokeEntityResp.length >= 1){
        if(revokeEntityResp[0]?.endDate){
            certificateStatus = SUSPENDED;
            msg = `Certificate is Suspended till ${revokeEntityResp[0]?.endDate}`
        }else{
            certificateStatus = REVOKED
            msg = `Certificate is Permanently Revoked`
        }
    }else{
        certificateStatus = VALID;
        msg = `certificate is Valid`;
    }
    return [certificateStatus,msg];
}

async function getEntity(entityId,filterType,certificateEntityType,token){
    const certificateFilter = {
        "filters": {
            [filterType]: {
                "eq": entityId
            }
        },
        "limit": 1,
        "offset": 0
    }
    return await sunbirdRegistryService.searchCertificate(certificateEntityType,certificateFilter,token);
}

async function deleteRevokeCertificate(req, res, kafkaProducer) {
    try {
        const token = req.header("Authorization");
        const filters = {
            "filters": {
                "previousCertificateId": {
                    "eq": req.params.revokedCertificateId
                },
                "schema": {
                    "eq": req.params.entityName
                }
            },
            "limit": 1,
            "offset": 0
        }

        let revokedVCResponse = await sunbirdRegistryService.searchCertificate("RevokedVC", filters, token)
        let revokedCertificateOsId = truncateShard(revokedVCResponse[0]?.osid);
        await kafkaProducer.connect();
        kafkaProducer.send({
            topic: certifyConstants.VC_REMOVE_SUSPENSION_TOPIC,
            messages: [
                { key: null, value: JSON.stringify({ revokedCertificateOsId: revokedCertificateOsId, token: token }) }
            ]
        });
        res.status(200).json({
            message: "Delete revoked certificate request sent successfully"
        });

    }
    catch (err) {
        console.log('ERROR : ', err?.response?.status || '');
        res.status(err?.response?.status || 500).json({
            message: err?.response?.data
        })
    }
}

module.exports = {
    createCertificate,
    getCertificate,
    updateCertificate,
    revokeCertificate,
    deleteRevokeCertificate,
    verifyCertificate
}


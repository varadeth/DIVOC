const config = require('./config');

const SUNBIRD_CERTIFICATE_URL = `${config.SUNBIRD_REGISTRY_URL}/api/v1/`;
const SVG_ACCEPT_HEADER = "image/svg+xml";
const IMAGE_RESPONSE_TYPE = 'svg';
const VC_CERTIFY_TOPIC = process.env.VC_CERTIFY_TOPIC || 'vc-certify';
const VC_REMOVE_SUSPENSION_TOPIC = process.env.VC_REMOVE_SUSPENSION_TOPIC || 'vc-remove-suspension';
const SUNBIRD_SSO_CLIENT = process.env.SUNBIRD_SSO_CLIENT || 'admin-api';
const SUNBIRD_SSO_ADMIN_CLIENT_SECRET = process.env.SUNBIRD_SSO_ADMIN_CLIENT_SECRET || '0358fa30-6014-4192-9551-7c61b15b774c';
const SUSPENSION_REMOVAL_SCHEDULE = process.env.SUSPENSION_REMOVAL_SCHEDULE || '0 0 * * *';
const REVOKED_ENTITY_TYPE = "RevokedVC";

module.exports = {
    SUNBIRD_CERTIFICATE_URL,
    SVG_ACCEPT_HEADER,
    IMAGE_RESPONSE_TYPE,
    VC_CERTIFY_TOPIC,
    VC_REMOVE_SUSPENSION_TOPIC,
    SUNBIRD_SSO_CLIENT,
    SUNBIRD_SSO_ADMIN_CLIENT_SECRET,
    SUSPENSION_REMOVAL_SCHEDULE,
    REVOKED_ENTITY_TYPE
}

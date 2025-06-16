package global_configs

// Header
const HEADER_LENGTH = 6
const STATUS_CODE_INDEX = 0
const OPERATION_CODE_INDEX = 1
const MESSAGE_LENGTH_INDEX = 2

// Hash
const HASH_LENGTH = 32

// Component Operation Codes
const USER_MESSAGE_OPERATION_CODE uint8 = 0

// Server Operation Codes
const REGISTER_COMPONENT_OPERATION_CODE uint8 = 255
const GET_COMPONENT_OPERATION_CODE uint8 = 254
const GET_COMPONENTS_OPERATION_CODE uint8 = 253

// Status Codes
const OK_STATUS uint8 = 0

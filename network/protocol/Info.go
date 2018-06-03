package protocol

const (
	PROTOCOL = 201
	VERSION  = "1.2.10.2"
)

var Protocols = map[int32]string{
	201: "1.2.10",
}

type PacketList map[PacketName]byte

type PacketName string

const (
	LOGIN_PACKET                         PacketName = "LoginPacket"
	PLAY_STATUS_PACKET                   PacketName = "PlayStatusPacket"
	SERVER_HANDSHAKE_PACKET              PacketName = "ServerHandshakePacket"
	CLIENT_HANDSHAKE_PACKET              PacketName = "ClientHandshakePacket"
	DISCONNECT_PACKET                    PacketName = "DisconnectPacket"
	RESOURCE_PACKS_INFO_PACKET           PacketName = "ResourcePackInfoPacket"
	RESOURCE_PACK_STACK_PACKET           PacketName = "ResourcePackStackPacket"
	RESOURCE_PACK_CLIENT_RESPONSE_PACKET PacketName = "ResourcePackClientResponsePacket"
	TEXT_PACKET                          PacketName = "TextPacket"
	SET_TIME_PACKET                      PacketName = "SetTimePacket"
	START_GAME_PACKET                    PacketName = "StartGamePacket"
	ADD_PLAYER_PACKET                    PacketName = "AddPlayerPacket"
	ADD_ENTITY_PACKET                    PacketName = "AddEntityPacket"
	REMOVE_ENTITY_PACKET                 PacketName = "RemoveEntityPacket"
	ADD_ITEM_ENTITY_PACKET               PacketName = "AddItemEntityPacket"
	ADD_HANGING_ENTITY_PACKET            PacketName = "AddHangingEntityPacket"
	TAKE_ITEM_ENTITY_PACKET              PacketName = "TakeItemEntityPacket"
	MOVE_ENTITY_PACKET                   PacketName = "MoveEntityPacket"
	MOVE_PLAYER_PACKET                   PacketName = "MovePlayerPacket"
	RIDER_JUMP_PACKET                    PacketName = "RiderJumpPacket"
	UPDATE_BLOCK_PACKET                  PacketName = "UpdateBlockPacket"
	ADD_PAINTING_PACKET                  PacketName = "AddPaintingPacket"
	EXPLODE_PACKET                       PacketName = "ExplodePacket"
	LEVEL_SOUND_EVENT_PACKET             PacketName = "LevelSoundEventPacket"
	LEVEL_EVENT_PACKET                   PacketName = "LevelEventPacket"
	BLOCK_EVENT_PACKET                   PacketName = "BlockEventPacket"
	ENTITY_EVENT_PACKET                  PacketName = "EntityEventPacket"
	MOB_EFFECT_PACKET                    PacketName = "MobEffectPacket"
	UPDATE_ATTRIBUTES_PACKET             PacketName = "UpdateAttributesPacket"
	INVENTORY_TRANSACTION_PACKET         PacketName = "InventoryTransactionPacket"
	MOB_EQUIPMENT_PACKET                 PacketName = "MobEquipmentPacket"
	MOB_ARMOR_EQUIPMENT_PACKET           PacketName = "MobArmorEquipmentPacket"
	INTERACT_PACKET                      PacketName = "InteractPacket"
	BLOCK_PICK_REQUEST_PACKET            PacketName = "BlockPickRequestPacket"
	ENTITY_PICK_REQUEST_PACKET           PacketName = "EntityPickRequestPacket"
	PLAYER_ACTION_PACKET                 PacketName = "PlayerActionPacket"
	ENTITY_FALL_PACKET                   PacketName = "EntityFallPacket"
	HURT_ARMOR_PACKET                    PacketName = "HurtArmorPacket"
	SET_ENTITY_DATA_PACKET               PacketName = "SetEntityDataPacket"
	SET_ENTITY_MOTION_PACKET             PacketName = "SetEntityMotionPacket"
	SET_ENTITY_LINK_PACKET               PacketName = "SetEntityLinkPacket"
	SET_HEALTH_PACKET                    PacketName = "SetHealthPacket"
	SET_SPAWN_POSITION_PACKET            PacketName = "SetSpawnPositionPacket"
	ANIMATE_PACKET                       PacketName = "AnimatePacket"
	RESPAWN_PACKET                       PacketName = "RespawnPacket"
	CONTAINER_OPEN_PACKET                PacketName = "ContainerOpenPacket"
	CONTAINER_CLOSE_PACKET               PacketName = "ContainerClosePacket"
	PLAYER_HOTBAR_PACKET                 PacketName = "PlayerHotbarPacket"
	INVENTORY_CONTENT_PACKET             PacketName = "InventoryContentPacket"
	INVENTORY_SLOT_PACKET                PacketName = "InventorySlotPacket"
	CONTAINER_SET_DATA_PACKET            PacketName = "ContainerSetDataPacket"
	CRAFTING_DATA_PACKET                 PacketName = "CraftingDataPacket"
	CRAFTING_EVENT_PACKET                PacketName = "CraftingEventPacket"
	GUI_DATA_PICK_ITEM_PACKET            PacketName = "GuiDataPickItemPacket"
	ADVENTURE_SETTINGS_PACKET            PacketName = "AdventureSettingsPacket"
	BLOCK_ENTITY_DATA_PACKET             PacketName = "BlockEntityDataPacket"
	PLAYER_INPUT_PACKET                  PacketName = "PlayerInputPacket"
	FULL_CHUNK_DATA_PACKET               PacketName = "FullChunkDataPacket"
	SET_COMMANDS_ENABLED_PACKET          PacketName = "SetCommandsEnabledPacket"
	SET_DIFFICULTY_PACKET                PacketName = "SetDifficultyPacket"
	CHANGE_DIMENSION_PACKET              PacketName = "ChangeDimensionPacket"
	SET_PLAYER_GAME_TYPE_PACKET          PacketName = "SetPlayerGameTypePacket"
	PLAYER_LIST_PACKET                   PacketName = "PlayerListPacket"
	SIMPLE_EVENT_PACKET                  PacketName = "SimpleEventPacket"
	EVENT_PACKET                         PacketName = "EventPacket"
	SPAWN_EXPERIENCE_ORB_PACKET          PacketName = "SpawnExperienceOrbPacket"
	CLIENTBOUND_MAP_ITEM_DATA_PACKET     PacketName = "ClientboundMapItemDataPacket"
	MAP_INFO_REQUEST_PACKET              PacketName = "MapInfoRequestPacket"
	REQUEST_CHUNK_RADIUS_PACKET          PacketName = "RequestChunkRadiusPacket"
	CHUNK_RADIUS_UPDATED_PACKET          PacketName = "ChunkRadiusUpdatedPacket"
	ITEM_FRAME_DROP_ITEM_PACKET          PacketName = "ItemFrameDropItemPacket"
	GAME_RULES_CHANGED_PACKET            PacketName = "GameRulesChangedPacket"
	CAMERA_PACKET                        PacketName = "CameraPacket"
	BOSS_EVENT_PACKET                    PacketName = "BossEventPacket"
	SHOW_CREDITS_PACKET                  PacketName = "ShowCreditsPacket"
	AVAILABLE_COMMANDS_PACKET            PacketName = "AvailableCommandsPacket"
	COMMAND_REQUEST_PACKET               PacketName = "CommandRequestPacket"
	COMMAND_BLOCK_UPDATE_PACKET          PacketName = "CommandBlockUpdatePacket"
	COMMAND_OUTPUT_PACKET                PacketName = "CommandOutputPacket"
	UPDATE_TRADE_PACKET                  PacketName = "UpdateTradePacket"
	UPDATE_EQUIP_PACKET                  PacketName = "UpdateEquipPacket"
	RESOURCE_PACK_DATA_INFO_PACKET       PacketName = "ResourcePackDataInfoPacket"
	RESOURCE_PACK_CHUNK_DATA_PACKET      PacketName = "ResourcePackChunkDataPacket"
	RESOURCE_PACK_CHUNK_REQUEST_PACKET   PacketName = "ResourcePackChunkRequestPacket"
	TRANSFER_PACKET                      PacketName = "TransferPacket"
	PLAY_SOUND_PACKET                    PacketName = "PlaySoundPacket"
	STOP_SOUND_PACKET                    PacketName = "StopSoundPacket"
	SET_TITLE_PACKET                     PacketName = "SetTitlePacket"
	ADD_BEHAVIOR_TREE_PACKET             PacketName = "AddBehaviorTreePacket"
	STRUCTURE_BLOCK_UPDATE_PACKET        PacketName = "StructureBlockUpdatePacket"
	SHOW_STORE_OFFER_PACKET              PacketName = "ShowStoreOfferPacket"
	PURCHASE_RECEIPT_PACKET              PacketName = "PurchaseReceiptPacket"
	PLAYER_SKIN_PACKET                   PacketName = "PlayerSkinPacket"
	SUB_CLIENT_LOGIN_PACKET              PacketName = "SubClientLoginPacket"
	W_S_CONNECT_PACKET                   PacketName = "WSConnectPacket"
	SET_LAST_HURT_BY_PACKET              PacketName = "SetLastHurtByPacket"
	BOOK_EDIT_PACKET                     PacketName = "BookEditPacket"
	NPC_REQUEST_PACKET                   PacketName = "NpcRequestPacket"
	PHOTO_TRANSFER_PACKET                PacketName = "PhotoTransferPacket"
	MODAL_FORM_REQUEST_PACKET            PacketName = "ModalFormRequestPacket"
	MODAL_FORM_RESPONSE_PACKET           PacketName = "ModalFormResponsePacket"
	SERVER_SETTINGS_REQUEST_PACKET       PacketName = "ServerSettingsRequestPacket"
	SERVER_SETTINGS_RESPONSE_PACKET      PacketName = "ServerSettingsResponsePacket"
	SHOW_PROFILE_PACKET                  PacketName = "ShowProfilePacket"
	SET_DEFAULT_GAME_TYPE_PACKET         PacketName = "SetDefaultGameTypePacket"
)

var packetId = PacketList{
	LOGIN_PACKET:                         0x01,
	PLAY_STATUS_PACKET:                   0x02,
	SERVER_HANDSHAKE_PACKET:              0x03,
	CLIENT_HANDSHAKE_PACKET:              0x04,
	DISCONNECT_PACKET:                    0x05,
	RESOURCE_PACKS_INFO_PACKET:           0x06,
	RESOURCE_PACK_STACK_PACKET:           0x07,
	RESOURCE_PACK_CLIENT_RESPONSE_PACKET: 0x08,
	TEXT_PACKET:                          0x09,
	SET_TIME_PACKET:                      0x0a,
	START_GAME_PACKET:                    0x0b,
	ADD_PLAYER_PACKET:                    0x0c,
	ADD_ENTITY_PACKET:                    0x0d,
	REMOVE_ENTITY_PACKET:                 0x0e,
	ADD_ITEM_ENTITY_PACKET:               0x0f,
	ADD_HANGING_ENTITY_PACKET:            0x10,
	TAKE_ITEM_ENTITY_PACKET:              0x11,
	MOVE_ENTITY_PACKET:                   0x12,
	MOVE_PLAYER_PACKET:                   0x13,
	RIDER_JUMP_PACKET:                    0x14,
	UPDATE_BLOCK_PACKET:                  0x15,
	ADD_PAINTING_PACKET:                  0x16,
	EXPLODE_PACKET:                       0x17,
	LEVEL_SOUND_EVENT_PACKET:             0x18,
	LEVEL_EVENT_PACKET:                   0x19,
	BLOCK_EVENT_PACKET:                   0x1a,
	ENTITY_EVENT_PACKET:                  0x1b,
	MOB_EFFECT_PACKET:                    0x1c,
	UPDATE_ATTRIBUTES_PACKET:             0x1d,
	INVENTORY_TRANSACTION_PACKET:         0x1e,
	MOB_EQUIPMENT_PACKET:                 0x1f,
	MOB_ARMOR_EQUIPMENT_PACKET:           0x20,
	INTERACT_PACKET:                      0x21,
	BLOCK_PICK_REQUEST_PACKET:            0x22,
	ENTITY_PICK_REQUEST_PACKET:           0x23,
	PLAYER_ACTION_PACKET:                 0x24,
	ENTITY_FALL_PACKET:                   0x25,
	HURT_ARMOR_PACKET:                    0x26,
	SET_ENTITY_DATA_PACKET:               0x27,
	SET_ENTITY_MOTION_PACKET:             0x28,
	SET_ENTITY_LINK_PACKET:               0x29,
	SET_HEALTH_PACKET:                    0x2a,
	SET_SPAWN_POSITION_PACKET:            0x2b,
	ANIMATE_PACKET:                       0x2c,
	RESPAWN_PACKET:                       0x2d,
	CONTAINER_OPEN_PACKET:                0x2e,
	CONTAINER_CLOSE_PACKET:               0x2f,
	PLAYER_HOTBAR_PACKET:                 0x30,
	INVENTORY_CONTENT_PACKET:             0x31,
	INVENTORY_SLOT_PACKET:                0x32,
	CONTAINER_SET_DATA_PACKET:            0x33,
	CRAFTING_DATA_PACKET:                 0x34,
	CRAFTING_EVENT_PACKET:                0x35,
	GUI_DATA_PICK_ITEM_PACKET:            0x36,
	ADVENTURE_SETTINGS_PACKET:            0x37,
	BLOCK_ENTITY_DATA_PACKET:             0x38,
	PLAYER_INPUT_PACKET:                  0x39,
	FULL_CHUNK_DATA_PACKET:               0x3a,
	SET_COMMANDS_ENABLED_PACKET:          0x3b,
	SET_DIFFICULTY_PACKET:                0x3c,
	CHANGE_DIMENSION_PACKET:              0x3d,
	SET_PLAYER_GAME_TYPE_PACKET:          0x3e,
	PLAYER_LIST_PACKET:                   0x3f,
	SIMPLE_EVENT_PACKET:                  0x40,
	EVENT_PACKET:                         0x41,
	SPAWN_EXPERIENCE_ORB_PACKET:          0x42,
	CLIENTBOUND_MAP_ITEM_DATA_PACKET:     0x43,
	MAP_INFO_REQUEST_PACKET:              0x44,
	REQUEST_CHUNK_RADIUS_PACKET:          0x45,
	CHUNK_RADIUS_UPDATED_PACKET:          0x46,
	ITEM_FRAME_DROP_ITEM_PACKET:          0x47,
	GAME_RULES_CHANGED_PACKET:            0x48,
	CAMERA_PACKET:                        0x49,
	BOSS_EVENT_PACKET:                    0x4a,
	SHOW_CREDITS_PACKET:                  0x4b,
	AVAILABLE_COMMANDS_PACKET:            0x4c,
	COMMAND_REQUEST_PACKET:               0x4d,
	COMMAND_BLOCK_UPDATE_PACKET:          0x4e,
	COMMAND_OUTPUT_PACKET:                0x4f,
	UPDATE_TRADE_PACKET:                  0x50,
	UPDATE_EQUIP_PACKET:                  0x51,
	RESOURCE_PACK_DATA_INFO_PACKET:       0x52,
	RESOURCE_PACK_CHUNK_DATA_PACKET:      0x53,
	RESOURCE_PACK_CHUNK_REQUEST_PACKET:   0x54,
	TRANSFER_PACKET:                      0x55,
	PLAY_SOUND_PACKET:                    0x56,
	STOP_SOUND_PACKET:                    0x57,
	SET_TITLE_PACKET:                     0x58,
	ADD_BEHAVIOR_TREE_PACKET:             0x59,
	STRUCTURE_BLOCK_UPDATE_PACKET:        0x5a,
	SHOW_STORE_OFFER_PACKET:              0x5b,
	PURCHASE_RECEIPT_PACKET:              0x5c,
	PLAYER_SKIN_PACKET:                   0x5d,
	SUB_CLIENT_LOGIN_PACKET:              0x5e,
	W_S_CONNECT_PACKET:                   0x5f,
	SET_LAST_HURT_BY_PACKET:              0x60,
	BOOK_EDIT_PACKET:                     0x61,
	NPC_REQUEST_PACKET:                   0x62,
	PHOTO_TRANSFER_PACKET:                0x63,
	MODAL_FORM_REQUEST_PACKET:            0x64,
	MODAL_FORM_RESPONSE_PACKET:           0x65,
	SERVER_SETTINGS_REQUEST_PACKET:       0x66,
	SERVER_SETTINGS_RESPONSE_PACKET:      0x67,
	SHOW_PROFILE_PACKET:                  0x68,
	SET_DEFAULT_GAME_TYPE_PACKET:         0x69,
}

func GetPacketId(name PacketName) byte {
	return packetId[name]
}

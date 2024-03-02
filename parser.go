package palworld_settings

import (
	"fmt"
	"strconv"
	"strings"
)

func Parse(settingString string) (*Setting, error) {
	optionSettingsIndex := strings.Index(settingString, "OptionSettings=(")

	if optionSettingsIndex == -1 {
		return nil, fmt.Errorf("OptionSettings not found in the settings file")
	}

	settingsStartIndex := optionSettingsIndex + len("OptionSettings=(")
	extractedSetting := settingString[settingsStartIndex:strings.LastIndex(settingString, ")")]

	return &Setting{
		Difficulty:                          getConfigValue(extractedSetting, "Difficulty"),
		DayTimeSpeedRate:                    getConfigValueFloat(extractedSetting, "DayTimeSpeedRate"),
		NightTimeSpeedRate:                  getConfigValueFloat(extractedSetting, "NightTimeSpeedRate"),
		ExpRate:                             getConfigValueFloat(extractedSetting, "ExpRate"),
		PalCaptureRate:                      getConfigValueFloat(extractedSetting, "PalCaptureRate"),
		PalSpawnNumRate:                     getConfigValueFloat(extractedSetting, "PalSpawnNumRate"),
		PalDamageRateAttack:                 getConfigValueFloat(extractedSetting, "PalDamageRateAttack"),
		PalDamageRateDefense:                getConfigValueFloat(extractedSetting, "PalDamageRateDefense"),
		PlayerDamageRateAttack:              getConfigValueFloat(extractedSetting, "PlayerDamageRateAttack"),
		PlayerDamageRateDefense:             getConfigValueFloat(extractedSetting, "PlayerDamageRateDefense"),
		PlayerStomachDecreaceRate:           getConfigValueFloat(extractedSetting, "PlayerStomachDecreaceRate"),
		PlayerStaminaDecreaceRate:           getConfigValueFloat(extractedSetting, "PlayerStaminaDecreaceRate"),
		PlayerAutoHPRegeneRate:              getConfigValueFloat(extractedSetting, "PlayerAutoHPRegeneRate"),
		PlayerAutoHpRegeneRateInSleep:       getConfigValueFloat(extractedSetting, "PlayerAutoHpRegeneRateInSleep"),
		PalStomachDecreaceRate:              getConfigValueFloat(extractedSetting, "PalStomachDecreaceRate"),
		PalStaminaDecreaceRate:              getConfigValueFloat(extractedSetting, "PalStaminaDecreaceRate"),
		PalAutoHPRegeneRate:                 getConfigValueFloat(extractedSetting, "PalAutoHPRegeneRate"),
		PalAutoHpRegeneRateInSleep:          getConfigValueFloat(extractedSetting, "PalAutoHpRegeneRateInSleep"),
		BuildObjectDamageRate:               getConfigValueFloat(extractedSetting, "BuildObjectDamageRate"),
		BuildObjectDeteriorationDamageRate:  getConfigValueFloat(extractedSetting, "BuildObjectDeteriorationDamageRate"),
		CollectionDropRate:                  getConfigValueFloat(extractedSetting, "CollectionDropRate"),
		CollectionObjectHpRate:              getConfigValueFloat(extractedSetting, "CollectionObjectHpRate"),
		CollectionObjectRespawnSpeedRate:    getConfigValueFloat(extractedSetting, "CollectionObjectRespawnSpeedRate"),
		EnemyDropItemRate:                   getConfigValueFloat(extractedSetting, "EnemyDropItemRate"),
		DeathPenalty:                        getConfigValue(extractedSetting, "DeathPenalty"),
		EnablePlayerToPlayerDamage:          getConfigValueBool(extractedSetting, "bEnablePlayerToPlayerDamage"),
		EnableFriendlyFire:                  getConfigValueBool(extractedSetting, "bEnableFriendlyFire"),
		EnableInvaderEnemy:                  getConfigValueBool(extractedSetting, "bEnableInvaderEnemy"),
		ActiveUNKO:                          getConfigValueBool(extractedSetting, "bActiveUNKO"),
		EnableAimAssistPad:                  getConfigValueBool(extractedSetting, "bEnableAimAssistPad"),
		EnableAimAssistKeyboard:             getConfigValueBool(extractedSetting, "bEnableAimAssistKeyboard"),
		DropItemMaxNum:                      getConfigValueInt(extractedSetting, "DropItemMaxNum"),
		DropItemMaxNum_UNKO:                 getConfigValueInt(extractedSetting, "DropItemMaxNum_UNKO"),
		BaseCampMaxNum:                      getConfigValueInt(extractedSetting, "BaseCampMaxNum"),
		BaseCampWorkerMaxNum:                getConfigValueInt(extractedSetting, "BaseCampWorkerMaxNum"),
		DropItemAliveMaxHours:               getConfigValueFloat(extractedSetting, "DropItemAliveMaxHours"),
		AutoResetGuildNoOnlinePlayers:       getConfigValueBool(extractedSetting, "bAutoResetGuildNoOnlinePlayers"),
		AutoResetGuildTimeNoOnlinePlayers:   getConfigValueFloat(extractedSetting, "AutoResetGuildTimeNoOnlinePlayers"),
		GuildPlayerMaxNum:                   getConfigValueInt(extractedSetting, "GuildPlayerMaxNum"),
		PalEggDefaultHatchingTime:           getConfigValueFloat(extractedSetting, "PalEggDefaultHatchingTime"),
		WorkSpeedRate:                       getConfigValueFloat(extractedSetting, "WorkSpeedRate"),
		IsMultiplay:                         getConfigValueBool(extractedSetting, "bIsMultiplay"),
		IsPvP:                               getConfigValueBool(extractedSetting, "bIsPvP"),
		CanPickupOtherGuildDeathPenaltyDrop: getConfigValueBool(extractedSetting, "bCanPickupOtherGuildDeathPenaltyDrop"),
		EnableNonLoginPenalty:               getConfigValueBool(extractedSetting, "bEnableNonLoginPenalty"),
		EnableFastTravel:                    getConfigValueBool(extractedSetting, "bEnableFastTravel"),
		IsStartLocationSelectByMap:          getConfigValueBool(extractedSetting, "bIsStartLocationSelectByMap"),
		ExistPlayerAfterLogout:              getConfigValueBool(extractedSetting, "bExistPlayerAfterLogout"),
		EnableDefenseOtherGuildPlayer:       getConfigValueBool(extractedSetting, "bEnableDefenseOtherGuildPlayer"),
		CoopPlayerMaxNum:                    getConfigValueInt(extractedSetting, "CoopPlayerMaxNum"),
		ServerPlayerMaxNum:                  getConfigValueInt(extractedSetting, "ServerPlayerMaxNum"),
		ServerName:                          getConfigValue(extractedSetting, "ServerName"),
		ServerDescription:                   getConfigValue(extractedSetting, "ServerDescription"),
		AdminPassword:                       getConfigValue(extractedSetting, "AdminPassword"),
		ServerPassword:                      getConfigValue(extractedSetting, "ServerPassword"),
		PublicPort:                          getConfigValueInt(extractedSetting, "PublicPort"),
		PublicIP:                            getConfigValue(extractedSetting, "PublicIP"),
		RCONEnabled:                         getConfigValueBool(extractedSetting, "RCONEnabled"),
		RCONPort:                            getConfigValueInt(extractedSetting, "RCONPort"),
		Region:                              getConfigValue(extractedSetting, "Region"),
		UseAuth:                             getConfigValueBool(extractedSetting, "bUseAuth"),
		BanListURL:                          getConfigValue(extractedSetting, "BanListURL"),
	}, nil
}

func getConfigValue(configString, configName string) string {
	configNameStartIndex := strings.Index(configString, configName)
	if configNameStartIndex == -1 {
		return ""
	}

	valueStart := strings.TrimPrefix(configString[strings.Index(configString, configName):], configName+"=")

	configLastIndex := strings.Index(valueStart, ",")
	if configLastIndex == -1 {
		return strings.Trim(valueStart, `"`)
	}
	return strings.Trim(valueStart[:strings.Index(valueStart, ",")], `"`)
}

func getConfigValueFloat(configString, configName string) float64 {
	value, err := strconv.ParseFloat(getConfigValue(configString, configName), 64)
	if err != nil {
		return 1
	}
	return value
}

func getConfigValueInt(configString, configName string) int {
	value, err := strconv.Atoi(getConfigValue(configString, configName))
	if err != nil {
		return 1
	}
	return value
}

func getConfigValueBool(configString, configName string) bool {
	return getConfigValue(configString, configName) == "true"
}

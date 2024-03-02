package palworld_settings

import "fmt"

func (s *Setting) ToString() string {
	configString := "[/Script/Pal.PalGameWorldSettings]\nOptionSettings=("
	configString += fmt.Sprintf("Difficulty=%s,", s.Difficulty)
	configString += fmt.Sprintf("DayTimeSpeedRate=%f,", s.DayTimeSpeedRate)
	configString += fmt.Sprintf("NightTimeSpeedRate=%f,", s.NightTimeSpeedRate)
	configString += fmt.Sprintf("ExpRate=%f,", s.ExpRate)
	configString += fmt.Sprintf("PalCaptureRate=%f,", s.PalCaptureRate)
	configString += fmt.Sprintf("PalSpawnNumRate=%f,", s.PalSpawnNumRate)
	configString += fmt.Sprintf("PalDamageRateAttack=%f,", s.PalDamageRateAttack)
	configString += fmt.Sprintf("PalDamageRateDefense=%f,", s.PalDamageRateDefense)
	configString += fmt.Sprintf("PlayerDamageRateAttack=%f,", s.PlayerDamageRateAttack)
	configString += fmt.Sprintf("PlayerDamageRateDefense=%f,", s.PlayerDamageRateDefense)
	configString += fmt.Sprintf("PlayerStomachDecreaceRate=%f,", s.PlayerStomachDecreaceRate)
	configString += fmt.Sprintf("PlayerStaminaDecreaceRate=%f,", s.PlayerStaminaDecreaceRate)
	configString += fmt.Sprintf("PlayerAutoHPRegeneRate=%f,", s.PlayerAutoHPRegeneRate)
	configString += fmt.Sprintf("PlayerAutoHpRegeneRateInSleep=%f,", s.PlayerAutoHpRegeneRateInSleep)
	configString += fmt.Sprintf("PalStomachDecreaceRate=%f,", s.PalStomachDecreaceRate)
	configString += fmt.Sprintf("PalStaminaDecreaceRate=%f,", s.PalStaminaDecreaceRate)
	configString += fmt.Sprintf("PalAutoHPRegeneRate=%f,", s.PalAutoHPRegeneRate)
	configString += fmt.Sprintf("PalAutoHpRegeneRateInSleep=%f,", s.PalAutoHpRegeneRateInSleep)
	configString += fmt.Sprintf("BuildObjectDamageRate=%f,", s.BuildObjectDamageRate)
	configString += fmt.Sprintf("BuildObjectDeteriorationDamageRate=%f,", s.BuildObjectDeteriorationDamageRate)
	configString += fmt.Sprintf("CollectionDropRate=%f,", s.CollectionDropRate)
	configString += fmt.Sprintf("CollectionObjectHpRate=%f,", s.CollectionObjectHpRate)
	configString += fmt.Sprintf("CollectionObjectRespawnSpeedRate=%f,", s.CollectionObjectRespawnSpeedRate)
	configString += fmt.Sprintf("EnemyDropItemRate=%f,", s.EnemyDropItemRate)
	configString += fmt.Sprintf("DeathPenalty=%s,", s.DeathPenalty)
	configString += fmt.Sprintf("bEnablePlayerToPlayerDamage=%s,", boolToString(s.EnablePlayerToPlayerDamage))
	configString += fmt.Sprintf("bEnableFriendlyFire=%s,", boolToString(s.EnableFriendlyFire))
	configString += fmt.Sprintf("bEnableInvaderEnemy=%s,", boolToString(s.EnableInvaderEnemy))
	configString += fmt.Sprintf("bActiveUNKO=%s,", boolToString(s.ActiveUNKO))
	configString += fmt.Sprintf("bEnableAimAssistPad=%s,", boolToString(s.EnableAimAssistPad))
	configString += fmt.Sprintf("bEnableAimAssistKeyboard=%s,", boolToString(s.EnableAimAssistKeyboard))
	configString += fmt.Sprintf("DropItemMaxNum=%d,", s.DropItemMaxNum)
	configString += fmt.Sprintf("DropItemMaxNum_UNKO=%d,", s.DropItemMaxNum_UNKO)
	configString += fmt.Sprintf("BaseCampMaxNum=%d,", s.BaseCampMaxNum)
	configString += fmt.Sprintf("BaseCampWorkerMaxNum=%d,", s.BaseCampWorkerMaxNum)
	configString += fmt.Sprintf("DropItemAliveMaxHours=%f,", s.DropItemAliveMaxHours)
	configString += fmt.Sprintf("bAutoResetGuildNoOnlinePlayers=%s,", boolToString(s.AutoResetGuildNoOnlinePlayers))
	configString += fmt.Sprintf("AutoResetGuildTimeNoOnlinePlayers=%f,", s.AutoResetGuildTimeNoOnlinePlayers)
	configString += fmt.Sprintf("GuildPlayerMaxNum=%d,", s.GuildPlayerMaxNum)
	configString += fmt.Sprintf("PalEggDefaultHatchingTime=%f,", s.PalEggDefaultHatchingTime)
	configString += fmt.Sprintf("WorkSpeedRate=%f,", s.WorkSpeedRate)
	configString += fmt.Sprintf("bIsMultiplay=%s,", boolToString(s.IsMultiplay))
	configString += fmt.Sprintf("bIsPvP=%s,", boolToString(s.IsPvP))
	configString += fmt.Sprintf("bCanPickupOtherGuildDeathPenaltyDrop=%s,", boolToString(s.CanPickupOtherGuildDeathPenaltyDrop))
	configString += fmt.Sprintf("bEnableNonLoginPenalty=%s,", boolToString(s.EnableNonLoginPenalty))
	configString += fmt.Sprintf("bEnableFastTravel=%s,", boolToString(s.EnableFastTravel))
	configString += fmt.Sprintf("bIsStartLocationSelectByMap=%s,", boolToString(s.IsStartLocationSelectByMap))
	configString += fmt.Sprintf("bExistPlayerAfterLogout=%s,", boolToString(s.ExistPlayerAfterLogout))
	configString += fmt.Sprintf("bEnableDefenseOtherGuildPlayer=%s,", boolToString(s.EnableDefenseOtherGuildPlayer))
	configString += fmt.Sprintf("CoopPlayerMaxNum=%d,", s.CoopPlayerMaxNum)
	configString += fmt.Sprintf("ServerPlayerMaxNum=%d,", s.ServerPlayerMaxNum)
	configString += fmt.Sprintf(`ServerName="%s",`, s.ServerName)
	configString += fmt.Sprintf(`ServerDescription="%s",`, s.ServerDescription)
	configString += fmt.Sprintf(`AdminPassword="%s",`, s.AdminPassword)
	configString += fmt.Sprintf(`ServerPassword="%s",`, s.ServerPassword)
	configString += fmt.Sprintf("PublicPort=%d,", s.PublicPort)
	configString += fmt.Sprintf(`PublicIP="%s",`, s.PublicIP)
	configString += fmt.Sprintf("RCONEnabled=%s,", boolToString(s.RCONEnabled))
	configString += fmt.Sprintf("RCONPort=%d,", s.RCONPort)
	configString += fmt.Sprintf(`Region="%s",`, s.Region)
	configString += fmt.Sprintf("bUseAuth=%s,", boolToString(s.UseAuth))
	configString += fmt.Sprintf(`BanListURL="%s"`, s.BanListURL)
	configString += ")"
	return configString
}

func boolToString(b bool) string {
	if b {
		return "True"
	}
	return "False"
}

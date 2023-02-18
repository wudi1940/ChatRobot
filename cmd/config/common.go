package config

const ProjectPath = "/root/MyProject/ChatRobot"

const MaxMsgContext = 10

const (
	InfoLogFileName  = "docs/logs/info_log.log"
	ErrorLogFileName = "docs/logs/err_log.log"
	GinLogFileName   = "docs/logs/gin_log.log"
)

func GetLanguageSelection() map[string][]string {
	return map[string][]string{
		"0": {"en-US", "zh-CN"},
		"1": {"zh-CN"},
		"2": {"en-US"},
		"3": {"fr-FR"},
		"4": {"de-DE"},
		"5": {"es-ES"},
	}
}

func GetLanguageSpeaker() map[string]string {
	return map[string]string{
		"0":       "",
		"1":       "zh-CN-YunxiNeural",
		"english": "en-US-JennyNeural",
		"french":  "fr-FR-DeniseNeural",
		"german":  "de-DE-ElkeNeural",
		"5":       "es-ES-AbrilNeural",
	}
}

const (
	AzureClientKey = ""
	AzureClientRegion
)

const OpenAIClientKey = ""

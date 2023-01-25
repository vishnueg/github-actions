package controller

import "initSetupScripts/utils"

func getJetsonUsecaseFolders(usecase string) (string, string, string) {
	var parent, folder, name string
	parent = utils.UC_FOLDER + "/" + utils.JETSON_EDGE

	switch usecase {
	case utils.FOOTBALL:
		folder = utils.FOOTBALL_UC
		name = utils.TAR_JETSON_FOOTBALL_UC
	case utils.REGULAR:
		folder = utils.REGULAR_UC
		name = utils.TAR_JETSON_REG_UC
	default:
		return "", "", ""
	}
	return parent, folder, name
}

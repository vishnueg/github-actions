package controller

import (
	"errors"
	utils "initSetupScripts/utils"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetRaspInit(c *gin.Context) {
	tarFile := utils.TAR_RASPI_INIT
	err := utils.CheckFileExistOrNot(tarFile)
	if errors.Is(err, os.ErrNotExist) {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"Message": "Tar file not found",
			})
		return
	} else if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"Message": "Some error while trying to fetch tar file",
			})
		return
	}
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+tarFile)
	c.Header("Content-Type", "application/octet-stream")
	c.File(tarFile)
}

func GetHostname(c *gin.Context) {
	log.Println("Fetching mac address from the url")
	macAddr := c.Param("macaddress")
	if macAddr == "" {
		log.Println("Mac-address is required ")
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Mac-address is required",
		})
		return
	}

	log.Println("Formatting mac address, replacing '-' with ':'")
	macAddr = strings.ReplaceAll(macAddr, "-", ":")
	log.Println("Fetching asset info from the DB")
	asset := utils.GetAssetInfo(macAddr)
	if asset == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error while trying to get asset info",
		})
		return
	} else if len(asset) == 0 {
		log.Println("No asset found with the following mac address: ", macAddr)
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "No asset found with the following mac address: " + macAddr,
		})
		return
	}

	log.Println("successful")
	c.JSON(
		http.StatusOK,
		gin.H{
			"hostname": asset[0].Asset_model,
		})
}

func GetJetsonInit(c *gin.Context) {
	tarFile := utils.TAR_JETSON_INIT
	err := utils.CheckFileExistOrNot(tarFile)
	if errors.Is(err, os.ErrNotExist) {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"Message": "Tar file not found",
			})
		return
	} else if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"Message": "Some error while trying to fetch tar file",
			})
		return
	}
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+tarFile)
	c.Header("Content-Type", "application/octet-stream")
	c.File(tarFile)
}

func GetRaspiUseCase(c *gin.Context) {
	log.Println("Fetching mac address from the url")

	macAddr := c.Param("macaddress")
	if macAddr == "" {
		log.Println("Mac-address is required ")
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Mac-address is required",
		})
		return
	}

	log.Println("Formatting mac address, replacing '-' with ':'")
	macAddr = strings.ReplaceAll(macAddr, "-", ":")

	log.Println("Fetching asset info from the DB")
	asset := utils.GetAssetInfo(macAddr)
	if asset == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error while trying to get asset info",
		})
		return
	} else if len(asset) == 0 {
		log.Println("No asset found with the following mac address: ", macAddr)
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "No asset found with the following mac address: " + macAddr,
		})
		return
	}

	tarFile := utils.TAR_RASPI_REG_UC
	if strings.EqualFold(strings.ToLower(asset[0].Description), strings.ToLower("Football")) {
		tarFile = utils.TAR_RASPI_FOOTBALL_UC
	}

	err := utils.CheckFileExistOrNot(tarFile)
	if errors.Is(err, os.ErrNotExist) {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"Message": "Tar file not found",
			})
		return
	} else if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"Message": "Some error while trying to fetch tar file",
			})
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+tarFile)
	c.Header("Content-Type", "application/octet-stream")
	c.File(tarFile)
}

func TarRaspInit(c *gin.Context) {
	parent := utils.INITPARENT
	folder := utils.RASPI_EDGE
	name := utils.TAR_RASPI_INIT
	log.Println("Creating a tar of following folder: ", parent+folder)
	if utils.CreateTarFile(parent, folder, name) != 0 {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"Message": "Error while trying to create a tar",
			})
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"Message": "Tar created successfully",
		})
}

func TarJetsonInit(c *gin.Context) {
	parent := utils.INITPARENT
	folder := utils.JETSON_EDGE
	name := utils.TAR_JETSON_INIT
	log.Println("Creating a tar of following folder: ", parent+folder)
	if utils.CreateTarFile(parent, folder, name) != 0 {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"Message": "Error while trying to create a tar",
			})
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"Message": "Tar created successfully",
		})
}

func TarRaspiRegularUseCase(c *gin.Context) {
	parent := utils.UC_FOLDER + "/" + utils.RASPI_EDGE
	folder := utils.REGULAR_UC
	name := utils.TAR_RASPI_REG_UC
	log.Println("Creating a tar of following folder: ", parent+folder)
	if utils.CreateTarFile(parent, folder, name) != 0 {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"Message": "Error while trying to create a tar",
			})
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"Message": "Tar created successfully",
		})
}

func TarRaspiFootballUseCase(c *gin.Context) {
	parent := utils.UC_FOLDER + "/" + utils.RASPI_EDGE
	folder := utils.FOOTBALL_UC
	name := utils.TAR_RASPI_FOOTBALL_UC
	log.Println("Creating a tar of following folder: ", parent+folder)
	if utils.CreateTarFile(parent, folder, name) != 0 {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"Message": "Error while trying to create a tar",
			})
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"Message": "Tar created successfully",
		})
}

func TarJetsonUseCase(c *gin.Context) {
	// parent := utils.UC_FOLDER + "/" + utils.JETSON_EDGE
	// folder := utils.REGULAR_UC
	usecase := c.Param("usecase")
	parent, folder, name := getJetsonUsecaseFolders(usecase)

	if parent == "" && folder == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"Message": "Invalid usecase provided",
			})
		return
	}
	log.Println("Creating a tar of following folder: ", parent+folder)
	if utils.CreateTarFile(parent, folder, name) != 0 {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"Message": "Error while trying to create a tar",
			})
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"Message": "Tar created successfully",
		})
}

func GetAssetUsecase(c *gin.Context) {
	log.Println("Fetching mac address from the url")
	macAddr := c.Param("macaddress")
	if macAddr == "" {
		log.Println("Mac-address is required ")
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Mac-address is required",
		})
		return
	}

	log.Println("Formatting mac address, replacing '-' with ':'")
	macAddr = strings.ReplaceAll(macAddr, "-", ":")
	log.Println("Fetching asset info from the DB")
	asset := utils.GetAssetInfo(macAddr)
	if asset == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Error while trying to get asset info",
		})
		return
	} else if len(asset) == 0 {
		log.Println("No asset found with the following mac address: ", macAddr)
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "No asset found with the following mac address: " + macAddr,
		})
		return
	}

	usecase := asset[0].Description
	c.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Usecase": usecase,
	})
}

func GetJetsonUseCase(c *gin.Context) {
	log.Println("Fetching mac address from the url")

	macAddr := c.Param("macaddress")
	if macAddr == "" {
		log.Println("Mac-address is required ")
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Mac-address is required",
		})
		return
	}

	log.Println("Formatting mac address, replacing '-' with ':'")
	macAddr = strings.ReplaceAll(macAddr, "-", ":")

	log.Println("Fetching asset info from the DB")
	asset := utils.GetAssetInfo(macAddr)
	if asset == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error while trying to get asset info",
		})
		return
	} else if len(asset) == 0 {
		log.Println("No asset found with the following mac address: ", macAddr)
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "No asset found with the following mac address: " + macAddr,
		})
		return
	}

	tarFile := utils.TAR_JETSON_REG_UC
	if strings.EqualFold(strings.ToLower(asset[0].Description), strings.ToLower("Football")) {
		tarFile = utils.TAR_JETSON_FOOTBALL_UC
	}

	err := utils.CheckFileExistOrNot(tarFile)
	if errors.Is(err, os.ErrNotExist) {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"Message": "Tar file not found",
			})
		return
	} else if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"Message": "Some error while trying to fetch tar file",
			})
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+tarFile)
	c.Header("Content-Type", "application/octet-stream")
	c.File(tarFile)
}

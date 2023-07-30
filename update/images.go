package update

import (
	"fmt"
	"os/exec"

	"github.com/dofusdude/ankabuffer"
	"github.com/dofusdude/api/utils"
)

func DownloadImagesLauncher(hashJson *ankabuffer.Manifest) error {

	fileNames := []HashFile{
		{Filename: "content/gfx/items/bitmap0.d2p", FriendlyName: "data/tmp/bitmaps_0.d2p"},
		{Filename: "content/gfx/items/bitmap0_1.d2p", FriendlyName: "data/tmp/bitmaps_1.d2p"},
		{Filename: "content/gfx/items/bitmap1.d2p", FriendlyName: "data/tmp/bitmaps_2.d2p"},
		{Filename: "content/gfx/items/bitmap1_1.d2p", FriendlyName: "data/tmp/bitmaps_3.d2p"},
		{Filename: "content/gfx/items/bitmap1_2.d2p", FriendlyName: "data/tmp/bitmaps_4.d2p"},
	}

	if err := DownloadUnpackFiles(hashJson, "main", fileNames, "", false); err != nil {
		return err
	}

	inPath := fmt.Sprintf("%s/data/tmp", utils.DockerMountDataPath)
	outPath := fmt.Sprintf("%s/data/img/item", utils.DockerMountDataPath)
	absConvertCmd := fmt.Sprintf("%s/PyDofus/%s_unpack.py", utils.DockerMountDataPath, "d2p")
	if err := exec.Command(utils.PythonPath, absConvertCmd, inPath, outPath).Run(); err != nil {
		return err
	}

	fileNames = []HashFile{
		{Filename: "content/gfx/items/vector0.d2p", FriendlyName: "data/tmp/vector/vector_0.d2p"},
		{Filename: "content/gfx/items/vector0_1.d2p", FriendlyName: "data/tmp/vector/vector_1.d2p"},
		{Filename: "content/gfx/items/vector1.d2p", FriendlyName: "data/tmp/vector/vector_2.d2p"},
		{Filename: "content/gfx/items/vector1_1.d2p", FriendlyName: "data/tmp/vector/vector_3.d2p"},
		{Filename: "content/gfx/items/vector1_2.d2p", FriendlyName: "data/tmp/vector/vector_4.d2p"},
	}

	if err := DownloadUnpackFiles(hashJson, "main", fileNames, "", false); err != nil {
		return err
	}

	inPath = fmt.Sprintf("%s/data/tmp/vector", utils.DockerMountDataPath)
	outPath = fmt.Sprintf("%s/data/vector/item", utils.DockerMountDataPath)
	if err := exec.Command(utils.PythonPath, absConvertCmd, inPath, outPath).Run(); err != nil {
		return err
	}

	return nil
}

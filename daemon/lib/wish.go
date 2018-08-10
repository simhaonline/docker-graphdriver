package lib

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type Wish struct {
	Id          int
	InputImage  int
	OutputImage int
	CvmfsRepo   string
}

type WishFriendly struct {
	Id         int
	InputId    int
	InputName  string
	OutputId   int
	OutputName string
	CvmfsRepo  string
}

func (d WishFriendly) PrintWish(machineFriendly, printHeader bool) {
	if machineFriendly {
		if printHeader {
			fmt.Println("id,input_image_id,input_image_name,cvmfs_repo,output_image_id,output_image_name")
		}
		fmt.Printf("%d,%d,%s,%s,%d,%s\n", d.Id, d.InputId, d.InputName, d.CvmfsRepo, d.OutputId, d.OutputName)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetHeader([]string{"Id", "Input Image Id",
			"Input Image Name", "CVMFS Repo",
			"Output Image Id", "Output Image Name"})
		table.Append([]string{strconv.Itoa(d.Id), strconv.Itoa(d.InputId),
			d.InputName, d.CvmfsRepo,
			strconv.Itoa(d.OutputId), d.OutputName})
		table.Render()
	}
}

func PrintMultipleWishes(wish []WishFriendly, machineFriendly, printHeader bool) {
	if machineFriendly {
		if printHeader {
			fmt.Println("id,input_image_id,input_image_name,cvmfs_repo,output_image_id,output_image_name")
		}
		for _, d := range wish {
			fmt.Printf("%d,%d,%s,%s,%d,%s\n", d.Id, d.InputId, d.InputName, d.CvmfsRepo, d.OutputId, d.OutputName)
		}
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		if printHeader {

			table.SetHeader([]string{"Id", "Input Image Id",
				"Input Image Name", "CVMFS Repo",
				"Output Image Id", "Output Image Name"})

		}
		for _, d := range wish {
			table.Append([]string{strconv.Itoa(d.Id), strconv.Itoa(d.InputId),
				d.InputName, d.CvmfsRepo,
				strconv.Itoa(d.OutputId), d.OutputName})
		}
		table.Render()
	}
}
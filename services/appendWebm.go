package services

import (
	"bytes"
	"errors"
)

func AppendTwoWebm(firstVidData, secdVidData []byte) ([]byte, error) {

	clusterStart1 := bytes.Index(firstVidData, []byte{0x1F, 0x43, 0xB6, 0x75})
    clusterStart2 := bytes.Index(secdVidData, []byte{0x1F, 0x43, 0xB6, 0x75})

    if clusterStart1 == -1 || clusterStart2 == -1 {
        return nil, errors.New("WebM data does not contain Cluster elements")
    }

    // Extract headers and data from both WebM files
    header1 := firstVidData[:clusterStart1]
    data1 := firstVidData[clusterStart1:]
    header2 := secdVidData[:clusterStart2]
    data2 := secdVidData[clusterStart2:]

    // Combine headers to create a new header
    newHeader := append(header1, header2...)

    // Append the new header and data from the second file to the first one
    newData := append(newHeader, data1...)
    newData = append(newData, data2...)

	return newData, nil
}
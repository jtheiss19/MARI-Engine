package gamestate

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/jtheiss19/MARI-Engine/elements"
	"github.com/jtheiss19/MARI-Engine/networking/mrp"
)

var ObjectMap = make(map[string]*elements.Element)
var blacklistedNames []string
var mu sync.Mutex

var chunkList []*Chunk

type Chunk struct {
	ChunkID   string
	ChunkData [][]*elements.Element
}

func CreateChunk() {
	ID := strconv.Itoa(len(chunkList))
	myNewChunk := &Chunk{ChunkID: ID}
	myNewChunk.ChunkData = [][]*elements.Element{}
	chunkList = append(chunkList, myNewChunk)
}

func GetEntireWorld() [][]*elements.Element {
	var masterMap = [][]*elements.Element{}

	for _, chunk := range chunkList {
		for _, layer := range chunk.ChunkData {
			masterMap = append(masterMap, layer)
		}
	}

	return masterMap
}

func GetEntireChunk(chunkID int) [][]*elements.Element {
	var masterMap = [][]*elements.Element{}

	masterMap = chunkList[chunkID].ChunkData

	return masterMap
}

func GetEntireChunkLayer(chunkID int, LayerID int) []*elements.Element {
	var masterMap = []*elements.Element{}

	masterMap = chunkList[chunkID].ChunkData[LayerID]

	return masterMap
}

func AddElemToChunk(elem *elements.Element, ChunkID int, PlaneToAdd int) {
	for {
		if ChunkID >= len(chunkList) {
			CreateChunk()
		} else {
			break
		}
	}

	for {
		if PlaneToAdd >= len(chunkList[ChunkID].ChunkData) {
			chunkList[ChunkID].ChunkData = append(chunkList[ChunkID].ChunkData, []*elements.Element{})
		} else {
			break
		}
	}

	elem.Layer = PlaneToAdd

	mu.Lock()
	for _, name := range blacklistedNames {
		if elem.UniqueName == name {
			mu.Unlock()
			return
		}
	}
	mu.Unlock()

	for _, unitElem := range chunkList[ChunkID].ChunkData[PlaneToAdd] {
		if unitElem.UniqueName == elem.UniqueName {
			mu.Lock()
			*unitElem = *elem
			mu.Unlock()
			return
		}
	}

	mu.Lock()
	chunkList[ChunkID].ChunkData[PlaneToAdd] = append(chunkList[ChunkID].ChunkData[PlaneToAdd], elem)
	mu.Unlock()

}

func GetObject(objectName string) *elements.Element {
	var returnElem *elements.Element = new(elements.Element)
	returnElem = ObjectMap[objectName].MakeCopy()
	return returnElem
}

func RemoveElem(badElem *elements.Element) {
	bytes, _ := json.Marshal(&badElem)

	myMRP := mrp.NewMRP([]byte("ELEM"), bytes, []byte("NIL"))

	for k, existing := range chunkList[0].ChunkData[badElem.Layer] {
		if badElem == nil || existing == nil {
			break
		}
		if badElem.UniqueName == existing.UniqueName {
			if k < len(chunkList[0].ChunkData[badElem.Layer]) {
				copy(chunkList[0].ChunkData[badElem.Layer][k:], chunkList[0].ChunkData[badElem.Layer][k+1:])
			}
			chunkList[0].ChunkData[badElem.Layer][len(chunkList[0].ChunkData[badElem.Layer])-1] = nil
			chunkList[0].ChunkData[badElem.Layer] = chunkList[0].ChunkData[badElem.Layer][:len(chunkList[0].ChunkData[badElem.Layer])-1]
		}
	}

	for _, conn := range connectionList {
		conn.Write(myMRP.MRPToByte())
	}
}

package gamestate

import (
	"encoding/json"
	"sync"

	"github.com/jtheiss19/MARI-Engine/elements"
	"github.com/jtheiss19/MARI-Engine/networking/mrp"
)

var ObjectMap = make(map[string]*elements.Element)
var blacklistedNames []string
var mu sync.Mutex

var MyChunk *Chunk

type Chunk struct {
	ChunkID   string
	ChunkData [][]*elements.Element
}

func GetEntireWorld() [][]*elements.Element {
	var masterMap = [][]*elements.Element{}

	for _, layer := range MyChunk.ChunkData {
		masterMap = append(masterMap, layer)
	}

	return masterMap
}

func GetEntireChunkLayer(LayerID int) []*elements.Element {
	var masterMap = []*elements.Element{}

	masterMap = MyChunk.ChunkData[LayerID]

	return masterMap
}

func AddElemToChunk(elem *elements.Element, PlaneToAdd int) {

	for {
		if PlaneToAdd >= len(MyChunk.ChunkData) {
			MyChunk.ChunkData = append(MyChunk.ChunkData, []*elements.Element{})
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

	for _, unitElem := range MyChunk.ChunkData[PlaneToAdd] {
		if unitElem.UniqueName == elem.UniqueName {
			mu.Lock()
			*unitElem = *elem
			mu.Unlock()
			return
		}
	}

	mu.Lock()
	MyChunk.ChunkData[PlaneToAdd] = append(MyChunk.ChunkData[PlaneToAdd], elem)
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

	for k, existing := range MyChunk.ChunkData[badElem.Layer] {
		if badElem == nil || existing == nil {
			break
		}
		if badElem.UniqueName == existing.UniqueName {
			if k < len(MyChunk.ChunkData[badElem.Layer]) {
				copy(MyChunk.ChunkData[badElem.Layer][k:], MyChunk.ChunkData[badElem.Layer][k+1:])
			}
			MyChunk.ChunkData[badElem.Layer][len(MyChunk.ChunkData[badElem.Layer])-1] = nil
			MyChunk.ChunkData[badElem.Layer] = MyChunk.ChunkData[badElem.Layer][:len(MyChunk.ChunkData[badElem.Layer])-1]
		}
	}

	for _, conn := range connectionList {
		conn.Write(myMRP.MRPToByte())
	}
}

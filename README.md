<div align="center">
  
![GitHub contributors](https://img.shields.io/github/contributors/jtheiss19/MARI-Engine)
![GitHub forks](https://img.shields.io/github/forks/jtheiss19/MARI-Engine?label=Forks)
![GitHub stars](https://img.shields.io/github/stars/jtheiss19/MARI-Engine?style=Stars)
![GitHub issues](https://img.shields.io/github/issues-raw/jtheiss19/MARI-Engine)
[![Go Report Card](https://goreportcard.com/badge/github.com/jtheiss19/MARI-Engine)](https://goreportcard.com/report/github.com/jtheiss19/MARI-Engine)

</div>

<h3 align="center">
    
  ![destroyer_img](https://cdn.discordapp.com/attachments/689340216284020812/689360748920832000/destroyer.png) MARI Engine ![destroyer_img](https://cdn.discordapp.com/attachments/689340216284020812/689360748920832000/destroyer.png)
  
An easy to implement game engine with built in multiplayer support
</h3>


<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
* [Technology Overview](#Technology)
  * [Elements](##Elements)
  * [Components](##Components)

<!-- ABOUT THE PROJECT -->
# About The Project 

The MARI engine is an easy to implement engine that can handle dynamic loading of maps via chunks and large scale multiplayer by leveraging the go languages natural ability to multi-thread with goroutines. 

# Technology Overview

The MARI Engine extends the [Ebiten](https://ebiten.org/) windowing/game engine to allow for programming by composition. It does this through the use of the Element and Component. The MARI Engine also can replicate any and all behavior from clients through its multiplayer framework. In fact, even when playing in a single player environment, the game will run identical to its server counterpart. This allows for easy testing and for easy networking. This scaling multiplayer network can handle large traffic volumes while maintaining deliverability through smart updating and replication culling.

## Elements

The Element is a structure found inside the MARI engine framework. Every object in the MARI engine is an Element. The Element is composed of a few default (almost universally used) data types. 

    type Element struct {
      XPos       float64
      YPos       float64
      Rotation   float64
      Active     bool
      UniqueName string
      ID         string
      Components []Component
      Same       bool
      Layer      int
    }

*   <b>XPos</b>      
Stores the world x position of the top left corner of the Element

*   <b>YPos</b>       
Stores the world y position of the top left corner of the Element

*   <b>Rotation</b>   
Stores the world rotation of the Element

*   <b>Active</b>     
Determines if the Element should be enabled or not. This is not for image culling but rather for setting an Element to be completely dormant. 

*   <b>UniqueName</b>       
The unique identifier for the Element. No two Elements may share a UniqueName during an online session. This is a nil value for all single player/non server related launches.

*   <b>ID</b>         
Determines the owner of the Element. Every player, even in single player, has one of these ID's associated with them. 

*   <b>Components</b>    
The extension of Element. The component slice is what adds all functionality to the Element.

## Components

    type Component interface {
      OnUpdate(xOffset float64, yOffset float64) error
      OnDraw(screen *ebiten.Image, xOffset float64, yOffset float64) error
      OnCheck(*Element) error
      OnMerge(Component) error
      OnUpdateServer() error
      MRP(finalElem *Element, conn net.Conn)
      SetContainer(*Element) error
      MakeCopy() Component
    }

A component is simply anything fitting the above interface. At first this may not seem impactful, but by having all objects in a game be the same structure allows for massive flexibility. 

Explanation by example is preferred. Pretend you have lots of trees and bushes in your game. Suddenly it comes upon you that your game must have fire physics. That fire should be able to spread from tree to tree and from tree to bush and bush to bush.

Before to accomplish this, you would have to write functions for each object you wanted to implement fire physics for and then assign each object its variables. This is tiresome especially if the physics needs to change. This would require changing every place you implemented code relating to fire physics. There are some ways around this dilemma with OOP design, but none are very elegant, nor do they scale very well as the object count grows.

Here is where MARI Engine shines. We create one component "FirePhysics" and write the code it needs to call for in the update or other respective functions. If we need variables custom to fire physics such as burnability and fire width, we can assign them to the custom component. Now we can add this component to anything and it will immediately implement our code. This kind of scalability is unparalleled. Not only this, but our code is now all in one location allowing for easy source control. 

## Declaring New Components 

## Installation



### Built With

* [Go](https://golang.org/)
* [Ebiten](https://ebiten.org/)
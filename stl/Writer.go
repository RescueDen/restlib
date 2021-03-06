// Copyright 2019 Reaction Engineering International. All rights reserved.
// Use of this source code is governed by the MIT license in the file LICENSE.txt.

package stl

import (
	"encoding/binary"
	"fmt"
	"io"
)

//Write the mesh, for now default to binary
func (mesh *Mesh) WriteMesh(w io.Writer) error {
	return mesh.WriteMeshBinary(w)
}

//Write the mesh out to the writer stream
func (mesh *Mesh) WriteMeshAscii(w io.Writer) error {

	//Now print each line
	io.WriteString(w, "solid "+mesh.Name+"\n")

	//For each element
	for _, ele := range mesh.Elements {
		//Get the normal
		norm := ele.getNormal()

		//Output Normal
		io.WriteString(w, "\tfacet normal ")
		printVertexString(w, &norm)
		io.WriteString(w, "\n\t\touter loop\n")

		//cycle over elements
		for _, node := range ele.Nodes {
			io.WriteString(w, "\t\t\tvertex ")
			printVertexString(w, &node)
			io.WriteString(w, "\n")
		}

		io.WriteString(w, "\t\tendloop\n")
		io.WriteString(w, "\tendfacet\n")

	}
	io.WriteString(w, "endsolid\n")

	return nil
}

//Write the mesh out to the writer stream as a binary file
func (mesh *Mesh) WriteMeshBinary(w io.Writer) error {

	//The header is always 80 bytes
	headerInfo := []byte(mesh.Name)

	//Current length
	l := len(headerInfo)

	//Resize the headerInfo
	if l > 80 {
		headerInfo = headerInfo[:80]
	} else if l < 80 {
		headerInfo = append(headerInfo, make([]byte, 80-l)...)

	}

	//Now write this to the file
	binary.Write(w, binary.LittleEndian, headerInfo)

	//Write out the number of elements
	binary.Write(w, binary.LittleEndian, uint32(len(mesh.Elements)))

	//For each element
	for _, ele := range mesh.Elements {
		//Get the normal
		norm := ele.getNormal()
		//Write it
		binary.Write(w, binary.LittleEndian, norm)

		//Now each vertex
		for _, node := range ele.Nodes {
			binary.Write(w, binary.LittleEndian, node)

		}

		//Output the attribute byte count, always zero
		binary.Write(w, binary.LittleEndian, uint16(0))

	}

	return nil
}

/**print vertex string**/
func printVertexString(w io.Writer, vertex *Vertex) {
	for _, num := range vertex {
		io.WriteString(w, fmt.Sprint(num)+" ")
	}
}

//GEt the stl as a pts file
func (mesh *Mesh) WriteUintahPts(w io.Writer) error {

	//March over each element and output teach pt
	for _, ele := range mesh.Elements {
		//Now each vertex
		for _, node := range ele.Nodes {
			printVertexString(w, &node)
			io.WriteString(w, "\n")
		}
	}
	return nil
}

//GEt the stl as a tri file
func (mesh *Mesh) WriteUintahTri(w io.Writer) error {

	//Assume each element was output in order with no shared nodes, one based index is used
	nodeNumber := 0

	//March over each element and output teach pt
	for _, ele := range mesh.Elements {
		//Now each vertex
		for range ele.Nodes {
			io.WriteString(w, fmt.Sprint(nodeNumber)+" ")

			//Bump the node number
			nodeNumber++
		}

		//Print an end line
		io.WriteString(w, "\n")
	}
	return nil
}

package main

var array1 = [5]int{1, 2, 3, 4, 5}
var array2 = [...]int{1, 2, 3, 4, 5}
var array3 = []int{1, 2, 3, 4, 5}
var array4 = [5]int{2: 3, 4: 5}
var array5 = []int{2: 3, 4: 5}

// Vector3D is 3d coordinate
type Vector3D [3]float32

var vec Vector3D

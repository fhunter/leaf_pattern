#!/usr/bin/env python

class Tree:
	def __init__(this, root):
		this.element=root
		this.branches=list()
	def walk(this):
		for i in this.branches:
			i.walk()
		this.action()
	def addbranch(this,element):
		this.branches.append(element)
	def action(this):
		print this.element

class Point:
	def __init__(this, x, y):
		this.point = complex(x,y)
	def __init__(this, cmplx):
		this.point = cmplx
		
		

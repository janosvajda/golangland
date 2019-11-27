Simple Factory
==============

Purpose
-------

https://en.wikipedia.org/wiki/Factory_(object-oriented_programming)


Usage
-----

.. code-block::  php
   :linenos:

	engine := GetEngine("airplane")
	Assamble(engine)
	Start(engine)

Run
----

go run DesignPatterns/Creational/SimpleFactory/SimpleFactory.go
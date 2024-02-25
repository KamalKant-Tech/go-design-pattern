### Builder Design Pattern

## What?
-- Builder desing pattern is creational design pattern which help us to create the object step by steps.
-- The pattern allows you to produce diff types and representation of objects using the same construction code.

## Why?
-- Some objects are easy to create with simple steps but some need special ceremonies to construct an object.
-- This is useful when you dont want to serve the unncessory information to your end user.

## How?
-- We can create diff struct(Builders) which going to do their respective jobs while constructing an object.
-- In this example we have construct three diff builder PersonBuilder, PersonAddressBuilder and PersonJobBuilder. Each builders are responsible to build an complex object.
-- For eg. PersonAddressBuilder will do the work for setting an object with person address, pin and other stuff.
-- The reason behind to have diff builder is that client can construct an object with their need if clients only worry abour whew a person address then he can construct an object with PersonAddressBuilder

## When do we need this?
-- When you don't need your end-user to provide millions of details in parameters of a constructor and you want to make his/her life easy by providing them with a simple way to create a complex object.


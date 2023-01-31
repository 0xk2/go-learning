# Thinking in Go

I am a Java person and OOP is the one I know from the very early day of my life so Go is a bit odd to me. You can use Go for OOP, but Go does not official support concepts such as inheritance or polymorphism.

## So, how to think in Go?

Oh, Inheritance is "trying to categorize things" while Go is more "tagging". Think of `struct` as defining the data structure and `interface` of defining the behavior of `something`, preferable call it `some code`.

Let's say we want to build a program for a car company who is going to sell a lot of different cars.

- The `Java way` is to build these classes with following inheritance tree
```
  Vehicle
   Car
    Suv
    Hatchback
    Sedan
    PickupTruck
    Convertibel
```
This kind of setup is useful as long as the inheritance tree reflect the truth the whole time. Unfortunately, it is not going to happen that way. For some odd reason, the original inheritane tree will go wrong (with the car dealer, it might not). This way enforce the will of the `architect` over the fact by naming things and conceptualize objects instead of describing the real world bit by bit.

However, there is nothing wrong with this way of thinking as long as your design is up-to-date with fact.

- The `Go way` is to find fact and describe it in code.
```
interface
  Movable

struct
  VehicleWithWheels
  VehicleWithSeats
```
With these basic interfaces and structs we can safely describe the world. There is no suv or sedan now; we can do it later by explain specifically what is a suv or sedan, how they work.

There is no overload, override in Go.
### Factory Design Pattern

## What?

- Factory design pattern in creational design pattern and one of the most commonly used pattern.
- Factory desing pattern is provide us an interface for creating the object but allows subclasses to alter the type or the objects that will be created.
- The main idea behind the Factory pattern is to create an abstraction layer over object creation. Instead of directly instantiating objects using their constructors, clients use a factory method or factory object to create objects.
- This provides flexibility by allowing the creation of different types of objects based on certain conditions or parameters.

## Why?

- <b>Encapsulation of Object Creation Logic</b>: The Factory pattern encapsulates the creation logic of objects within a separate component, such as a factory method or factory class. This separation of concerns helps to keep object creation logic isolated from the client code, promoting cleaner and more maintainable code.

- <b>Abstraction over Object Creation</b>: By providing a factory method or factory class, the pattern abstracts the process of object creation. Clients interact with the factory interface rather than directly instantiating objects using constructors. This abstraction allows for greater flexibility in object creation, such as creating different types of objects based on runtime conditions or configuration.

- <b>Centralized Object Creation</b>: Factories provide a centralized place for managing object creation logic. This centralization can simplify maintenance and debugging, as all changes related to object creation are confined to the factory implementation.

- <b>Decoupling of Client and Concrete Classes</b>: The Factory pattern promotes loose coupling between client code and concrete classes. Clients rely on the factory interface rather than concrete class implementations, reducing dependencies and making the code more adaptable to changes.

- <b>Support for Dependency Injection</b>: Factories can be used to implement dependency injection, where objects are passed to clients rather than being created directly. This approach facilitates testing, as dependencies can be easily mocked or replaced with alternative implementations.

## Where?

- <b>Creating Objects Dynamically</b>: When you want to create objects dynamically at runtime based on parameters or conditions, the Factory pattern can be useful. For example, you might have different types of database connectors (MySQL, PostgreSQL, MongoDB), and the choice of connector depends on configuration settings or user input.

- <b>Managing Complex Object Creation Logic</b>: If object creation involves complex logic, such as initialization with default values, validation, or resource management, the Factory pattern can encapsulate this logic in a dedicated factory class or method. This simplifies client code by abstracting away the complexity of object creation.

```
Overall, the Factory design pattern is valuable in scenarios where you need to abstract object creation, encapsulate creation logic, promote flexibility, and adhere to design principles. It provides a structured approach to managing object instantiation, improving code organization, and facilitating future enhancements and maintenance.

```


## How?

```
We define a Product interface with a PrintName method that all concrete products must implement.

We create two concrete product types, Electronics and Furniture, each of which implements the PrintName method.

We implement a CreateProduct factory function that takes a categoty argument and returns a Product based on the provided type. This function acts as a factory for creating products.

In the main function, we use the CreateProduct factory function to create instances of Electronics and Furniture, and we call the PrintName method on each product.

```
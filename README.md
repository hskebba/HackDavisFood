# HackDavisFood

This document contains a technical overview of the HackDavis FoodService. This document contains information related to specification on how we are building the feature and any technical considerations.

###### tags: `README` `food_service`

[TOC]

## Problem Statement

* Yolo Food Bank needs a way to schedule food pickups with available drivers
* Yolo Food Bank wants to report on nutritional information of items in inventory
* Pantry Users need to be informed when new items arrive


## Solution

We plan to directly connect suppliers, Food Bank delivery drivers, and Pantry users so all parties have real time, up to date information.

When a [Supplier] has an [Order] that needs to be picked up by a [Driver], the [Supplier] will initate the request by sending a message of PICKUP to a Food Bank number in Twilio. From the Twilio API, we will start a converstation with the [Supplier], gathering information like the location, pickup times, and a description of the order. Once the [Supplier] has competed the pickup request, the [Order] will be available for a driver to cliam.

A driver can claim an open order from the [Orders UI]. Once the order is claimed, it will be assigned to that driver. When the order is delivered to the Food Bank, the delivery driver will being updating the [Inventory] with the new [Products]. Using a web app, the driver will update the [Inventory] for each [Product]. As an optional step, if the [Product] has listed nutritional information the driver can take a picture of the information, parsed with image recognition and it will be stored alongside the [Product] and be displayed to users. 

When the driver has completed updating the [Inventory], the order will be marked as Complete. At that time, users subscribed to recive updates will be notfied via text message of the newly available [Products]. Given more time, we would make users subscriptions filterable, so that users could only be notified on certin products becoming available or when a product exceeds a certin quantity. 

## Services

### Orders 

#### Overview

| **Service Name**           | orders                                         |
| -------------------------- | ---------------------------------------------- |
| **Description**            | Accept new orders and update in progress orders |
| **Ports**                  | 50052                                          |
| **External Access Method** | GRPC                                           |

### Inventory 

#### Overview

| **Service Name**           | inventory                                                                      |
| -------------------------- | ------------------------------------------------------------------------------ |
| **Description**            | Communicate with the inventory ui so users can see what inventory is available |
| **Ports**                  | 50052                                                                          |
| **External Access Method** | GRPC                                                                           |

### Data_api 

#### Overview

| **Service Name**           | data_api                                 |
| -------------------------- | ---------------------------------------- |
| **Description**            | Handle CRUD opperations for all services |
| **Ports**                  | 50052                                    |
| **External Access Method** | GRPC                                     |

### Food_frontend 

#### Overview

| **Service Name**           | food_frontend                                                                        |
| -------------------------- | ------------------------------------------------------------------------------------ |
| **Description**            | All interfacing with data. Allows drivers to view orders, users to see inventory |
| **External Access Method** | https                                                                                |

#### Required Environment Vairables

| ENV Name      | Env Value | Default         | Required |
| ------------- | --------- | --------------- | -------- |
| COCKROACH_URI |           | cockroach:26257 |          |

## Concepts

### Supplier

The supplier is an entity that is trying to have excess or unwanted food delivered to the Food Bank. The supplier is resonsible for initiating delivery requests.

### Order

An order is a request from a [Supplier] to pickup food to deliver to the Food Bank. 

### OrderStatus

* OPEN: the Supplier has made a request for a new Order 
* ACCEPTED: a Driver has accepted an order and is enroute 
* PICKED_UP: a Driver has loaded the order on the truck and is enroute to the Food Bank
* PROCESSING: a Driver has reached the Food Bank and is updating the inventory
* COMPLETE: the Driver has finished unloading the order and updating the inventory

### Driver

The drive is responsible for picking up food from the supplier, initaiting their own [Order] pickups, and updating the inventory.

### Products

Products are items the Food Bank stocks. Ex apples, bananas, chicken would be considered products. 

### NutritionalData

The nutrition data for a [Product]

### Inventory

Each item in the Food Bank is an individual entry in the Inventory. Ex if there are 3 apples, there would be 3 rows of [Product] apple in the Inventory. This is done to track which Order each item came from.

### Users

End users who are browsing the Food Bank for new items

### Subscriptions

Users can chose to be notified of new [Products]. This creates a Subscription for the [User] and [Product]. 

## Schemas

### CockroachDB

#### foodbank.users

| Column | Type  | Primary Key | Default | Constraints | Foreign Key | Description              |
| ------ | ----- | ----------- | ------- | ----------- | ----------- | ------------------------ |
| id     | UUID4 | X           | uuid4() |             |             | Unique Id of the User    |
| phone  | TEXT  |             |         | NOT NULL    |             | Phone Number of the user |
| name   | TEXT  |             |         | NOT NULL    |             | Name of the user         |

#### foodbank.subscriptions

| Column     | Type  | Primary Key | Default | Constraints | Foreign Key | Description                   |
| ---------- | ----- | ----------- | ------- | ----------- | ----------- | ----------------------------- |
| id         | UUID4 | X           |         | NOT NULL    |             | Unique ID of the Subscription |
| user_id    | UUID4 |             |         |             | users.id    | Foreign Key lookup to User    |
| product_id | UUID4 |             |         |             | products.id | Foreign Key lookup to Product |

#### foodbank.driver

| Column | Type  | Primary Key | Default | Constraints | Foreign Key | Description             |
| ------ | ----- | ----------- | ------- | ----------- | ----------- | ----------------------- |
| id     | UUID4 | X           |         | NOT NULL    |             | Unique ID of the Driver |
| name   | TEXT  |             |         | NOT NULL    |             | name of the Driver      |

#### foodbank.orders

| Column      | Type     | Primary Key | Default | Constraints | Foreign Key  | Description                                            |
| ----------- | -------- | ----------- | ------- | ----------- | ------------ | ------------------------------------------------------ |
| id          | TEXT     | X           |         | NOT NULL    |              | Unique ID of the Order                                 |
| description | TEXT     |             |         |             |              | Supplier's description of the Order                    |
| pickup_at   | DATETIME |             |         |             |              | Time of Order pickup                                   |
| location    | TEXT     |             |         |             |              | Place where the Order will be picked up                |
| supplier_id | UUID4    |             |         | NOT NULL    | suppliers.id | Foreign Key lookup to Supplier                         |
| status      | TEXT     |             | OPEN    |             |              | Current status of Order                                |
| image_id    | UUID4    |             |         |             | images.id    | Foreign Key lookup to Image                            |
| updated_at  | DATETIME |             |         |             |              | Time of last update made to Order                      |
| created_at  | DATETIME |             | Now()   |             |              | Time the Order was created                             |
| driver_id   | UUID4    |             |         |             | drivers.id   | Foreign Key lookup to the driver who claimed the order |

#### foodbank.suppliers

| Column | Type  | Primary Key | Default | Constraints      | Foreign Key | Description                  |
| ------ | ----- | ----------- | ------- | ---------------- | ----------- | ---------------------------- |
| id     | UUID4 | X           | uuid4() |                  |             | Unique ID of the Supplier    |
| phone  | TEXT  |             |         | NOT NULL, Unique |             | Phone Number of the Supplier |
| name   | TEXT  |             |         | NOT NULL         |             | Name of the Supplier         |

#### foodbank.images

| Column | Type  | Primary Key | Default | Constraints | Foreign Key | Description            |
| ------ | ----- | ----------- | ------- | ----------- | ----------- | ---------------------- |
| id     | TEXT  | X           | uuid4() |             |             | Unique ID of the Image |
| data   | BYTES |             |         |             |             | Bytes of the Image     | 


#### foodbank.inventory

| Column       | Type     | Primary Key | Default | Constraints | Foreign Key | Description                   |
| ------------ | -------- | ----------- | ------- | ----------- | ----------- | ----------------------------- |
| id           | UUID4()  | X           | uuid4() |             |             | Unique ID of the Inventory    |
| product_id   | UUID4()  |             |         | NOT NULL    | products.id | Foreign Key lookup to Product |
| order_id     | UUID4()  |             |         | NOT NULL    | orders.id   | Foreign Key lookup to Order   |
| delivered_at | DATETIME |             |         |             |             | Time order was completed      |
| expired_at   | DATETIME |             |         |             |             | Time item will expire         |


#### foodbank.nutritional_data

| Column     | Type    | Primary Key | Default | Constraints | Foreign Key | Description                       |
| ---------- | ------- | ----------- | ------- | ----------- | ----------- | --------------------------------- |
| id         | UUID4() | X           | uuid4() |             |             | Unique ID of the Nutritional Data |
| product_id | UUID4() |             |         | NOT NULL    | products.id | Foreign Key lookup to Product     |
| fat        | FLOAT   |             |         | NOT NULL    |             | Amount of Fat in Product          |
| calories   | FLOAT   |             |         | NOT NULL    |             | Amount of Calories in Product     |
| sodium     | FLOAT   |             |         | NOT NULL    |             | Amount of Sodium in Product       |
| carbs      | FLOAT   |             |         | NOT NULL    |             | Amount of Carbs in Product        |
| protein    | FLOAT   |             |         | NOT NULL    |             | Amount of Protein in Product      |


#### foodbank.products

| Column     | Type    | Primary Key | Default | Constraints | Foreign Key | Description                 |
| ---------- | ------- | ----------- | ------- | ----------- | ----------- | --------------------------- |
| id         | UUID4() | X           | uuid4() |             |             | Unique ID of the Products   |
| name       | TEXT    |             |         | NOT NULL    |             | Name of Product             |
| image_id   | UUID4() |             |         |             | images.id   | Foreign Key lookup to Image |
| tag_set_id |         |             |         |             |             |                             |

#### foodbank.tag_sets

| Column | Type    | Primary Key | Default | Constraints | Foreign Key | Description |
| ------ | ------- | ----------- | ------- | ----------- | ----------- | ----------- |
| id     | UUID4() | X           |         |             |             |             |
| tag_id | UUID4() | X           |         |             |             |             |


#### foodbank.tags

| Column | Type    | Primary Key | Default | Constraints | Foreign Key | Description |
| ------ | ------- | ----------- | ------- | ----------- | ----------- | ----------- |
| id     | UUID4() | X           |         |             |             |             |
| name   |         |             |         | UNIQUE      |             |             |

## Endpoints

### orders

#### GetOrders

``` proto
// Order tracks a request from a Supplier to pickup food and deliver to the Food Bank
message Order {
    // Id unique identifer for an order
    string id = 1;
    // Description user supplied info of the order
    string description = 2;
    // PickupAt user suplied datetime for the requested pickup time
    google.protobuf.Timestamp pickup_at = 3;
    // Location address of the pickup
    string location = 4;
    // SupplierId id of the supplier requesting the pickup
    string supplier_id = 5;
    // Status indicates what stage the Order is in
    string status = 6;
    // Image bytes of the image of the Order provided by the supplier
    repeated byte image = 7;
    // UpdatedAt when the Order was last updated
    google.protobuf.Timestamp updated_at = 8;
    // CreatedAt when the order was created
    google.protobuf.Timestamp created_at = 9;
}

// NewOrderRequest request from a Supplier for a new Order
// For this request, supplier_id is not required as it will be determined by the phone_number
// Additionaly, supplier_name is only required for new suppliers
message NewOrderRequest {
    // Order details of the order
    Order order = 1;
    // PhoneNumber number the request was submitted from
    string phone_number = 2;
    // SupplierName name of the supplier
    string supplier_name = 3;
}

// NewOrderResponse empty response on success
message NewOrderResponse {}

// NewOrder process a new Order from a supplier
rpc NewOrder (NewOrderRequest) returns (NewOrderResponse) {}
```

#### ClaimOrder

``` protobuf
// ClaimOrderRequest 
message ClaimOrderRequest {
    string driver_id = 1;
    string order_id = 2;
}

// ClaimOrderResponse
message ClaimOrderResponse {}

// ClaimORder assigns and order to a driver
rpc ClaimOrder (ClaimOrderRequest) returns (ClaimOrderResponse) {}
```

#### CancelOrder

``` protobuf


#### ProcessOrder

#### CompleteOrder

#### GetProducts

### inventory

#### InsertProduct

#### UpdateInventory

### data_api

#### UpsertOrder

#### InsertProduct

#### UpdateInventory


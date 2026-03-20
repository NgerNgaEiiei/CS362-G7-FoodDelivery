# Food Delivery Platform
---

## Step 1: Data Design

### 1. Entities และ Attributes (ข้อมูลที่จำเป็น)
   
#### Customer
| Attribute | Type | Description | 
| :--- | :--- | :--- | 
| CustomerID (PK) | Integer | รหัสผู้ใช้ | 
| CustomerName | String | ชื่อผู้ใช้ | 
| CustomerPhone | VARCHAR | เบอร์โทรผู้ใช้ | 
| CustomerAddress | String | ที่อยู่ผู้ใช้ | 

#### Restaurant
| Attribute | Type | Description | 
| :--- | :--- | :--- | 
| RestaurantID (PK) | Integer | รหัสร้านอาหาร | 
| RestaurantName | String | ชื่อร้านอาหาร | 
| RestaurantLocation | Geo | ที่ตั้งของร้าน | 

> **หมายเหตุ:** `Geo` คือ ชนิดข้อมูลที่ใช้สำหรับเก็บ "พิกัดทางภูมิศาสตร์" เพื่อใช้คำนวณระยะทางจัดส่ง

#### Order
| Attribute | Type | Description | 
| :--- | :--- | :--- | 
| orderId | Integer | id ของรายการอาหารนั้น ๆ | 
| orderDate | DATE | วันที่ที่ทำการสั่งอาหาร | 
| totalPrice | Integer | ราคารวมทั้งหมดในการสั่งซื้อ | 
| status | String | สถานะของการสั่งซื้อ | 
| DeliveryAddr | Geo | ที่อยู่จัดส่งของลูกค้า | 

#### OrderItem
| Attribute | Type | Description | 
| :--- | :--- | :--- | 
| quantity | Integer | จำนวนของรายการอาหารที่สั่ง | 
| subTotal | Integer | ราคารวมย่อยของแต่ละรายการอาหาร | 
| specialInstructions | Text | รายละเอียดเพิ่มเติม | 

---

## Step 2: Architectural Mapping

### Layered Architecture

#### Controller Layer
* `AuthController`
* `CustomerController`
* `RestaurantController`
* `OrderController`
* `CartController`
* `RiderController`
    
#### Service Layer
* `AuthSerVice`
* `OrderService`
* `CartService`
* `RestaurantService`
* `DeliveryService`
   
#### Repository Layer
* `CustomerRepository`
* `OrderRepository`
* `RestaurantRepository`
* `CartRepository`
* `MenuRepository`

---

## Folder Structure Example
```bash
FoodDeliveryPlatform/                   
├── models/
│   └── order.go
│   └── restaurant.go
│   └── menu.go
├── repositories/
│   └── rider_repository.go  
│   └── order_repository.go 
│   └── menu_repository.go  
├── services/
│   └── restaurant_service.go    
│   └── rider_service.go    
│   └── order_service.go    
├── handlers/
│   └── rider_handler.go     
│   └── restaurant_handler.go    
│   └── order_handler.go
└── main.go   
```

## Step 3: Interface & Controller Contract
### Task 1: Controller (API Specification)
### 1) Add Item to Cart
- **Endpoint:** `POST /cart/items`
- **Description:** เพิ่มรายการอาหารลงในตะกร้าของลูกค้า

#### Request Body
```json
{
  "customerId": 1,
  "foodItemId": 101,
  "quantity": 2,
  "specialInstructions": "ไม่ใส่ผัก"
}
```

#### Success Response
```json
{
  "message": "Item added to cart successfully",
  "cartId": 5001,
  "item": {
    "foodItemId": 101,
    "quantity": 2,
    "unitPrice": 60,
    "specialInstructions": "ไม่ใส่ผัก"
  }
}
```

#### Error Response
```json
{
  "message": "Food item not found"
}
```

---

### 2) Get Cart by Customer
- **Endpoint:** `GET /cart/{customerId}`
- **Description:** ดูรายการอาหารทั้งหมดในตะกร้าของลูกค้า

#### Success Response
```json
{
  "cartId": 5001,
  "customerId": 1,
  "items": [
    {
      "orderItemId": 1,
      "foodItemId": 101,
      "foodName": "ข้าวผัดกุ้ง",
      "quantity": 2,
      "unitPrice": 60,
      "subTotal": 120,
      "specialInstructions": "ไม่ใส่ผัก"
    }
  ],
  "totalPrice": 120
}
```

#### Error Response
```json
{
  "message": "Cart not found"
}
```

---

### 3) Update Cart Item
- **Endpoint:** `PUT /cart/items/{orderItemId}`
- **Description:** แก้ไขจำนวนหรือรายละเอียดเพิ่มเติมของรายการอาหารในตะกร้า

#### Request Body
```json
{
  "quantity": 3,
  "specialInstructions": "เผ็ดน้อย"
}
```

#### Success Response
```json
{
  "message": "Cart item updated successfully"
}
```

#### Error Response
```json
{
  "message": "Cart item not found"
}
```

---

### 4) Remove Cart Item
- **Endpoint:** `DELETE /cart/items/{orderItemId}`
- **Description:** ลบรายการอาหารออกจากตะกร้า

#### Success Response
```json
{
  "message": "Cart item removed successfully"
}
```

#### Error Response
```json
{
  "message": "Cart item not found"
}
```

---

### 5) Create Order
- **Endpoint:** `POST /orders`
- **Description:** ยืนยันคำสั่งซื้อจากรายการในตะกร้าของลูกค้า

#### Request Body
```json
{
  "customerId": 1,
  "restaurantId": 10,
  "deliveryAddr": "89/12 Khlong Luang, Pathum Thani"
}
```

#### Success Response
```json
{
  "message": "Order created successfully",
  "orderId": 9001,
  "status": "PENDING",
  "totalPrice": 120
}
```

#### Error Response
```json
{
  "message": "Cart is empty"
}
```

---

### 6) Get Order Detail
- **Endpoint:** `GET /orders/{orderId}`
- **Description:** ดูรายละเอียดคำสั่งซื้อที่ถูกสร้างแล้ว

#### Success Response
```json
{
  "orderId": 9001,
  "orderDate": "2026-03-19",
  "customerId": 1,
  "restaurantId": 10,
  "status": "PENDING",
  "deliveryAddr": "89/12 Khlong Luang, Pathum Thani",
  "items": [
    {
      "orderItemId": 1,
      "foodItemId": 101,
      "foodName": "ข้าวผัดกุ้ง",
      "quantity": 2,
      "subTotal": 120,
      "specialInstructions": "ไม่ใส่ผัก"
    }
  ],
  "totalPrice": 120
}
```

#### Error Response
```json
{
  "message": "Order not found"
}
```
### Task 2: Interface
### 1) CartService Interface
- **Description:** จัดการตะกร้าสินค้าของลูกค้า (เพิ่ม, แก้ไข, ลบ, ดูรายการ)
``` Go
   type CartService interface {
    AddItemToCart(customerID int, foodItemID int, quantity int, specialInstructions string) (Cart, error)
    GetCartByCustomer(customerID int) (Cart, error)
    UpdateCartItem(orderItemID int, quantity int, specialInstructions string) error
    RemoveCartItem(orderItemID int) error
}
```

### 2) Order Service Interface
- **Description:** จัดการคำสั่งซื้อ (สร้างคำสั่งซื้อ และดูรายละเอียดคำสั่งซื้อ)
``` Go
   type OrderService interface {
    CreateOrder(customerID int, restaurantID int, deliveryAddr string) (Order, error)
    GetOrderDetail(orderID int) (Order, error)
}
```

### Task 3: Refinement (Entity Logic)
### 1) Cart Entity Logic
``` Go
   type Cart struct {
       CartID     int
       CustomerID int
       Items      []OrderItem
       TotalPrice int
   }
   
   // คำนวณราคารวมของตะกร้า
   func (c *Cart) CalculateTotal() {
       total := 0
       for _, item := range c.Items {
           total += item.CalculateSubTotal()
       }
       c.TotalPrice = total
   }
   
   // ตรวจสอบว่าตะกร้าว่างหรือไม่
   func (c *Cart) IsEmpty() bool {
       return len(c.Items) == 0
   }
```
- **CalculateTotal()** ใช้คำนวณราคารวมของสินค้าในตะกร้าจาก OrderItem ทั้งหมด เพื่อให้ได้ราคารวมล่าสุดก่อนแสดงผลหรือสร้างคำสั่งซื้อ
- **IsEmpty()** ใช้ตรวจสอบว่าตะกร้าว่างหรือไม่ เพื่อป้องกันการสร้างคำสั่งซื้อจากตะกร้าที่ไม่มีสินค้า

### 2) Order Entity Logic
``` Go
   type Order struct {
       OrderID      int
       OrderDate    string
       CustomerID   int
       RestaurantID int
       Status       string
       DeliveryAddr string
       Items        []OrderItem
       TotalPrice   int
   }
   
   // คำนวณราคารวมของคำสั่งซื้อ
   func (o *Order) CalculateTotal() {
       total := 0
       for _, item := range o.Items {
           total += item.CalculateSubTotal()
       }
       o.TotalPrice = total
   }
   
   // ตรวจสอบความถูกต้องของคำสั่งซื้อ
   func (o *Order) IsValid() bool {
       return o.CustomerID > 0 &&
              o.RestaurantID > 0 &&
              len(o.Items) > 0
   }
   
```
- **CalculateTotal()** ใช้คำนวณราคารวมของคำสั่งซื้อจากรายการอาหารทั้งหมดใน Order
- **IsValid()** ใช้ตรวจสอบว่าข้อมูลคำสั่งซื้อครบถ้วน เช่น มีลูกค้า ร้านอาหาร และรายการอาหาร ก่อนบันทึกลงระบบ

### 3) OrderItem Entity Logic
``` Go
   type OrderItem struct {
       OrderItemID         int
       FoodItemID          int
       Quantity            int
       UnitPrice           int
       SpecialInstructions string
   }
   
   // คำนวณราคาย่อยของแต่ละรายการ
   func (oi *OrderItem) CalculateSubTotal() int {
       return oi.UnitPrice * oi.Quantity
   }
   
   // ตรวจสอบข้อมูลรายการอาหาร
   func (oi *OrderItem) IsValid() bool {
       return oi.Quantity > 0 && oi.UnitPrice >= 0
   }
```
- **CalculateSubTotal()** ใช้คำนวณราคารวมย่อยของแต่ละรายการอาหารจากจำนวนและราคาต่อหน่วย
- **IsValid()** ใช้ตรวจสอบความถูกต้องของข้อมูล เช่น จำนวนต้องมากกว่า 0 และราคาต้องไม่ติดลบ

### 4) Customer Entity Logic
``` Go
   type Customer struct {
       CustomerID      int
       CustomerName    string
       CustomerPhone   string
       CustomerAddress string
   }
   
   // ตรวจสอบข้อมูลลูกค้า
   func (c *Customer) IsValid() bool {
       return c.CustomerName != "" &&
              c.CustomerPhone != "" &&
              c.CustomerAddress != ""
   }
```
- **IsValid()** ใช้ตรวจสอบว่าข้อมูลลูกค้ามีความครบถ้วน เช่น ชื่อ เบอร์โทร และที่อยู่ ก่อนนำไปใช้งานในระบบ

### 5) Restaurant Entity Logic
``` Go
   type Restaurant struct {
       RestaurantID       int
       RestaurantName     string
       RestaurantLocation string
   }
   
   // ตรวจสอบข้อมูลร้านอาหาร
   func (r *Restaurant) IsValid() bool {
       return r.RestaurantName != "" &&
              r.RestaurantLocation != ""
   }
```
- **IsValid()** ใช้ตรวจสอบว่าข้อมูลร้านอาหารถูกต้อง เช่น มีชื่อร้านและตำแหน่งที่ตั้ง ก่อนใช้งานในระบบ

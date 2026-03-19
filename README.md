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
* `DoctorController`
* `AppointmentController`
* `RecordController`
   
#### Service Layer
* `AuthSerVice`
* `DoctorService`
* `AppointmentService`
* `ZoomService`
* `MedicalRecordService`
   
#### Repository Layer
* `UserRepository`
* `DoctorDetailRepository`
* `TimeSlotRepository`
* `AppointmentRepository`
* `MedicalRecordRepository`

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
- **Endpoint:** `POST /api/cart/items`
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
- **Endpoint:** `GET /api/cart/{customerId}`
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
- **Endpoint:** `PUT /api/cart/items/{orderItemId}`
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
- **Endpoint:** `DELETE /api/cart/items/{orderItemId}`
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
- **Endpoint:** `POST /api/orders`
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
- **Endpoint:** `GET /api/orders/{orderId}`
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

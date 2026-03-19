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

## Step 3: Interface & Controller Contract

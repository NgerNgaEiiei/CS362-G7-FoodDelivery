# Food Delivery Platform
## Step 1: Data Design
1. Entities และ Attributes (ข้อมูลที่จำเป็น)
   
Customer

| Attribute | Type | Description | 
|---|---|---|
| CustomerID (PK) | Integer | รหัสผู้ใช้ | 
| CustomerName | String | ชื่อผู้ใช้ | 
| CustomerPhone | VARCHAR | เบอร์โทรผู้ใช้ | 
| CustomerAddress | String | ที่อยู่ผู้ใช้ | 

Restaurant

| Attribute | Type | Description | 
|---|---|---|
| RestaurantID (PK) | Integer |รหัสร้านอาหาร | 
| RestaurantName | String |ชื่อร้านอาหาร | 
| RestaurantLocation | Geo | ที่ตั้งของร้าน | 

หมายเหตุ: Geo คือ ชนิดข้อมูลที่ใช้สำหรับเก็บ "พิกัดทางภูมิศาสตร์" เพื่อใช้คำนวณระยะทางจัดส่ง


---
## Step 2: Architectural Mapping
---
## Step 3: Interface & Controller Contract

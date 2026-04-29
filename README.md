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
| CustomerAddress | Geo | ที่อยู่ผู้ใช้ | 

#### Restaurant
| Attribute | Type | Description | 
| :--- | :--- | :--- | 
| RestaurantID (PK) | Integer | รหัสร้านอาหาร | 
| RestaurantName | String | ชื่อร้านอาหาร | 
| RestaurantPhone | String | เบอร์โทรร้าน | 
| RestaurantLocation | Geo | ที่ตั้งของร้าน | 

> **หมายเหตุ:** `Geo` คือ ชนิดข้อมูลที่ใช้สำหรับเก็บ "พิกัดทางภูมิศาสตร์" เพื่อใช้คำนวณระยะทางจัดส่ง

#### Rider
| Attribute | Type | Description | 
| :--- | :--- | :--- | 
| RiderID (PK) | Integer | รหัสไรเดอร์ | 
| RiderName | String | ชื่อไรเดอร์ | 
| RiderPhone | VARCHAR | เบอร์โทรไรเดอร์ | 
| CurrentLocation | Geo | พิกัดปัจจุบันของไรเดอร์ ใช้จับคู่กับออเดอร์ที่ใกล้ที่สุด |
| Status | VARCHAR | สถานะ: AVAILABLE / BUSY / SELECTED / GOING_OFFLINE / OFFLINE |

#### FoodItem (Menu)
| Attribute | Type | Description |
| :--- | :--- | :--- |
| FoodItemID (PK) | Integer | รหัสอาหาร |
| RestaurantID (FK) | Integer | อ้างอิงร้านอาหาร |
| FoodName | String | ชื่อเมนู |
| Price | Integer | ราคา |
| Description | Text | รายละเอียด |
| IsAvailable | Boolean | สถานะว่ายังขายอยู่ไหม |

#### Cart
| Attribute | Type | Description | 
| :--- | :--- | :--- | 
| CartID (PK) | Integer | รหัสตะกร้า | 
| CustomerID (FK) | Integer | อ้างอิงถึงลูกค้าเจ้าของตะกร้า | 
| RestaurantID (FK) | Integer | อ้างอิงร้านอาหาร ใช้ validate ว่า item ทุกอันมาจากร้านเดียวกัน | 
| UpdatedAt | Timestamp | เวลาที่ตะกร้าถูกอัปเดตล่าสุด |

#### CartItem
| Attribute | Type | Description | 
| :--- | :--- | :--- | 
| CartItemID (PK) | Integer | รหัสรายการอาหารในตะกร้า | 
| CartID (FK) | Integer | อ้างอิงถึงตะกร้า | 
| FoodItemID (FK) | Integer | อ้างอิงเมนูอาหาร | 
| Quantity | Integer | จำนวนที่สั่ง เปลี่ยนได้ตลอดก่อนกด checkout | 

#### Order
| Attribute | Type | Description | 
| :--- | :--- | :--- | 
| OrderID (PK) | Integer | รหัสคำสั่งซื้อ | 
| CustomerID (FK) | Integer | อ้างอิงถึงลูกค้า | 
| RestaurantID (FK) | Integer | อ้างอิงถึงร้านอาหาร | 
| RiderID (FK) | Integer | อ้างอิงถึงไรเดอร์ (NULL ได้ถ้าระบบยังหา rider ไม่เจอ) | 
| Status | VARCHAR | สถานะของการสั่งซื้อ | 
| TotalPrice | Integer | ราคารวมทั้งหมด snapshot ไว้ตอนสั่ง | 
| DeliveryAddress | Geo | ที่อยู่จัดส่ง อาจไม่ใช่บ้านลูกค้าก็ได้ | 
| CreatedAt | TIMESTAMP | เวลาที่กด confirm ออเดอร์ | 
| CustomerGracePeriodEnd | TIMESTAMP | deadline ที่ลูกค้ายังยกเลิกได้ หลังเวลานี้ยกเลิกไม่ได้แล้ว | 
| RestaurantGracePeriodEnd | TIMESTAMP | deadline ที่ร้านยังปฏิเสธออเดอร์ได้ ประมาณ 10 วินาทีหลังรับออเดอร์ | 

#### OrderItem
| Attribute | Type | Description | 
| :--- | :--- | :--- | 
| OrderItemID (PK) | Integer | รหัสรายการอาหารในออเดอร์ | 
| OrderID (FK) | Integer | อ้างอิงถึงคำสั่งซื้อ | 
| FoodItemID (FK) | Integer | อ้างอิงเมนูอาหาร |
| Quantity | Integer | จำนวนที่สั่ง | 
| Subtotal | Integer | ราคา ณ ตอนสั่ง × quantity ไม่เปลี่ยนแม้ร้านจะแก้ราคาทีหลัง | 
| SpecialInstructions | Text | คำขอพิเศษเฉพาะ item นี้ เช่น "ไม่เอาผัก", "เผ็ดน้อย" |

## Step 2: Architectural Mapping

### Layered Architecture

#### Controller Layer
* `CartController`
* `OrderController`
* `MenuController`
* `RiderController`
    
#### Service Layer
* `CartService`
* `OrderService`
* `MenuService`
* `RiderService`
* `RiderAssignmentService`
* `NotificationService`
   
#### Repository Layer
* `CartRepository`
* `OrderRepository`
* `MenuRepository`
* `RestaurantRepository`
* `RiderRepository`

---

## Folder Structure Example
```bash
CS362-G7-FoodDelivery/
├── controllers/
│   ├── cart_controller.go
│   ├── order_controller.go
│   ├── menu_controller.go
│   └── rider_controller.go
│
├── services/
│   ├── cart_service.go
│   ├── order_service.go
│   ├── menu_service.go
│   ├── rider_service.go
│   ├── rider_assignment_service.go
│   └── notification_service.go
│
├── repositories/
│   ├── cart_repository.go
│   ├── order_repository.go
│   ├── menu_repository.go
│   ├── restaurant_repository.go
│   └── rider_repository.go
│
├── models/
│   ├── cart.go
│   ├── order.go
│   ├── order_item.go
│   ├── customer.go
│   ├── restaurant.go
│   ├── food_item.go
│   ├── rider.go
│   ├── assignment.go
│   ├── notification.go
│   └── geo.go
│
└── README.md 
```

## Step 3: Interface & Controller Contract
### Cart Service
#### Interface
```go
type CartService interface {
    AddItem(userID string, item CartItem) error
    RemoveItem(userID string, foodItemID string) error
    UpdateQuantity(userID string, foodItemID string, quantity int) error
    GetCart(userID string) (*Cart, error)
    ClearCart(userID string) error
}
 
type CartRepository interface {
    Save(cart *Cart) error
    FindByUserID(userID string) (*Cart, error)
    DeleteByUserID(userID string) error
}
```

#### Controller Contract
```
GET    /cart
  Response: { "cart_id": "abc", "items": [...], "total": 150.00 }
 
POST   /cart/items
  Request:  { "food_item_id": "123", "quantity": 2 }
  Response: { "cart_id": "abc", "total": 150.00 }
 
PUT    /cart/items/:food_item_id
  Request:  { "quantity": 3 }
  Response: { "cart_id": "abc", "total": 225.00 }
 
DELETE /cart/items/:food_item_id
  Response: { "message": "item removed" }
 
DELETE /cart
  Response: { "message": "cart cleared" }
```

---

### Order Service

#### Interface
```go
type OrderService interface {
    CreateOrder(userID string, cartID string) (*Order, error)
    GetOrder(orderID string) (*Order, error)
    GetOrdersByUser(userID string) ([]*Order, error)
    CancelOrder(orderID string) error
    UpdateOrderStatus(orderID string, status OrderStatus) error
}
 
type OrderRepository interface {
    Save(order *Order) error
    FindByID(orderID string) (*Order, error)
    FindByUserID(userID string) ([]*Order, error)
    UpdateStatus(orderID string, status OrderStatus) error
}
```

#### Controller Contrct
```
POST   /orders
  Request:  { "cart_id": "abc", "address": "123 Main St" }
  Response: { "order_id": "xyz", "status": "pending", "total": 150.00 }
 
GET    /orders/:order_id
  Response: { "order_id": "xyz", "status": "delivering", "items": [...] }
 
GET    /orders
  Response: [ { "order_id": "xyz", "status": "completed", ... }, ... ]
 
DELETE /orders/:order_id
  Response: { "message": "order cancelled" }
```

---

### Menu / Restaurant Service

#### Interface
```go
type MenuService interface {
    GetRestaurants() ([]*Restaurant, error)
    GetRestaurantByID(restaurantID string) (*Restaurant, error)
    GetMenuByRestaurant(restaurantID string) ([]*MenuItem, error)
    AddMenuItem(restaurantID string, item MenuItem) error
    UpdateMenuItem(itemID string, item MenuItem) error
    DeleteMenuItem(itemID string) error
}
 
type MenuRepository interface {
    FindAllRestaurants() ([]*Restaurant, error)
    FindRestaurantByID(restaurantID string) (*Restaurant, error)
    FindMenuByRestaurantID(restaurantID string) ([]*MenuItem, error)
    SaveMenuItem(item *MenuItem) error
    UpdateMenuItem(item *MenuItem) error
    DeleteMenuItem(itemID string) error
}
```

#### Controller Contract
```
GET    /restaurants
  Response: [ { "restaurant_id": "r1", "name": "...", "rating": 4.5 }, ... ]
 
GET    /restaurants/:restaurant_id
  Response: { "restaurant_id": "r1", "name": "...", "menu": [...] }
 
GET    /restaurants/:restaurant_id/menu
  Response: [ { "item_id": "m1", "name": "...", "price": 80.00 }, ... ]
 
POST   /restaurants/:restaurant_id/menu
  Request:  { "name": "Pad Thai", "price": 80.00, "category": "main" }
  Response: { "item_id": "m1", "name": "Pad Thai", "price": 80.00 }
 
PUT    /restaurants/:restaurant_id/menu/:item_id
  Request:  { "price": 90.00 }
  Response: { "item_id": "m1", "price": 90.00 }
 
DELETE /restaurants/:restaurant_id/menu/:item_id
  Response: { "message": "item deleted" }
```
 
---

### Rider Service

#### Interface
```go
type RiderService interface {
    GetRiderByID(riderID string) (*Rider, error)
    UpdateLocation(riderID string, lat float64, lng float64) error
    SetAvailability(riderID string, available bool) error
    GetActiveOrders(riderID string) ([]*Order, error)
    UpdateOrderStatus(riderID string, orderID string, status OrderStatus) error
}
 
type RiderRepository interface {
    FindByID(riderID string) (*Rider, error)
    UpdateLocation(riderID string, lat float64, lng float64) error
    UpdateAvailability(riderID string, available bool) error
    FindActiveOrdersByRiderID(riderID string) ([]*Order, error)
}
```

#### Controller Contract
```
GET    /riders/:rider_id
  Response: { "rider_id": "r1", "name": "...", "available": true }
 
PUT    /riders/:rider_id/location
  Request:  { "lat": 13.7563, "lng": 100.5018 }
  Response: { "message": "location updated" }
 
PUT    /riders/:rider_id/availability
  Request:  { "available": true }
  Response: { "message": "availability updated" }
 
GET    /riders/:rider_id/orders
  Response: [ { "order_id": "xyz", "status": "picking_up", ... }, ... ]
 
PUT    /riders/:rider_id/orders/:order_id/status
  Request:  { "status": "delivered" }
  Response: { "message": "status updated" }
```
 
---

### Rider Assignment Service
#### Interface
```go
type RiderAssignmentService interface {
    AssignRider(order *Order) (*Assignment, error)
    FindNearestRider(lat float64, lng float64) (*Rider, error)
    IsRiderAvailable(riderID string) (bool, error)
    ReleaseRider(riderID string, orderID string) error
}
```
 
#### Event Contract
```
Consumes: order.created
  Payload: { "order_id": "xyz", "restaurant_lat": 13.75, "restaurant_lng": 100.50 }
 
Publishes: rider.assigned
  Payload: { "order_id": "xyz", "rider_id": "r1", "estimated_pickup": "10min" }
```
 
---

### Notification Service
> ทำงานแบบ async ผ่าน Kafka — ไม่มี HTTP endpoint

#### Interface
```go
type NotificationService interface {
    SendPushNotification(userID string, message Notification) error
    NotifyCustomer(orderID string, event string) error
    NotifyRestaurant(orderID string, event string) error
    NotifyRider(riderID string, event string) error
}
```
 
#### Event Contract
```
Consumes: rider.assigned
  Action: แจ้ง rider ว่ามีงานใหม่
 
Consumes: order.status_updated
  Action: แจ้ง customer ว่า order status เปลี่ยน
 
Consumes: restaurant.refused
  Action: แจ้ง customer ว่าร้านปฏิเสธ order
```


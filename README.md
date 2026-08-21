# 🪙 API de Sistema de Recompensas y Puntos

API REST desarrollada en Go para la gestión de acumulación, consulta y redención de puntos por compras realizadas por los clientes.

---

## 📌 Reglas de Negocio

* **Acumulación:** Por cada **$1.000 COP** en compras, el cliente obtiene **1 punto**.
* **Redención:** Cada **1 punto** equivale a **$100 COP** al momento de redimir.

---

## 🚀 Requerimientos Funcionales

1. **Registrar una compra:** Permite ingresar el monto de una transacción efectuada por el cliente.
2. **Calcular puntos obtenidos:** Aplica la conversión automática del monto gastado a puntos ganados.
3. **Acumular puntos:** Suma los nuevos puntos obtenidos al saldo histórico del cliente por múltiples compras.
4. **Consultar saldo:** Retorna el total acumulado de puntos disponibles para redimir.
5. **Redimir puntos:** Permite al cliente intercambiar sus puntos disponibles por su equivalente en dinero/descuento.
6. **Validación de saldo insuficiente:** Si el cliente intenta redimir más puntos de los que posee, el sistema retorna un error indicando la insuficiencia de saldo.

---

## 🛠️ Endpoints de la API 

| Método | Endpoint | Descripción |
| :--- | :--- | :--- |
| `POST` | `/api/v1/purchases` | Registra una compra y calcula los puntos generados. |
| `GET` | `/api/v1/customers/{id}/points` | Consulta el saldo total de puntos de un cliente. |
| `POST` | `/api/v1/customers/{id}/redeem` | Redime una cantidad específica de puntos del cliente. |

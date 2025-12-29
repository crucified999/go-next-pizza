import React from "react"

type OrderPaymentsProps = {
  payments: string
}

export const OrderPayments: React.FC<OrderPaymentsProps> = ({ payments }) => {
  return (
    <div>
      <h3>Способы оплаты</h3>
      
    </div>
  )
}
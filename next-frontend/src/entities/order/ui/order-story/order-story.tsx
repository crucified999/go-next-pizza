import React from "react";
import { Order } from "../../model/types"
import { formatDateTimeExact } from "../../lib/utils";

type OrderStoryProps = {
  orders: Order[];
}

export const OrderStory: React.FC<OrderStoryProps> = ({ orders }) => {
  return orders.length > 0 ? (
    <div className="mb-20">
      <h3 className="font-bold text-2xl mb-10">История заказов</h3>
      <table className="flex-col">
        <thead>
          <tr className="text-black/50 border-b-1">
            <th className="text-left p-5">№</th>
            <th className="text-left p-5">Время заказа</th>
            <th className="text-left p-5">Сумма</th>
          </tr>
        </thead>
        <tbody className="font-bold">
          {
            orders.map((order) => 
              <tr key={order.id} className="border-b-1">
                <td className="p-5 text-left">{order.id}</td>
                <td className="p-5">{formatDateTimeExact(order.createdAt)}</td>
                <td className="p-5 text-center">{order.totalPrice} ₽</td>
              </tr>
            )
          }
        </tbody>
      </table>
    </div>
  ) : 
    (
      <div>
        <h3 className="font-bold text-2xl mb-15x">У вас пока нет заказов, но это можно исправить!</h3>
      </div>
    )
}
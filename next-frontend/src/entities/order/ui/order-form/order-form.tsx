import { User } from "@/entities/user/model/types";
import { DataInput } from "@/shared/ui/data-input";
import React, { SyntheticEvent } from "react";

type OrderFormProps = {
  user: User;
};

export const OrderForm: React.FC<OrderFormProps> = ({ user }) => {

  const handleSubmit = (e: SyntheticEvent) => {
    e.preventDefault();
  }

  return (
    <form onSubmit={handleSubmit} className="flex flex-col gap-5 mt-10 w-full">
      <DataInput label="Имя" value={user.name} canChange inputClassname="text-black" />
      <DataInput label="Номер телефона" value={user.phone} canChange={false} />
      <DataInput label="Адрес доставки" value={""} canChange inputClassname="w-100 text-black" />
    </form>
  );
};

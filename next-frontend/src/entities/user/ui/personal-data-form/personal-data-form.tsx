import { useAppSelector } from "@/app/store";
import { DataInput } from "@/shared/ui/data-input";
import { SyntheticEvent } from "react";
import { changeEmail, changeName } from "../../lib/api";
import { formatPhoneFromE164 } from "@/shared/lib/utils";

export const PersonalDataForm = () => {
  // const phone = sessionStorage?.getItem("phone-number");

  const phone = useAppSelector((state) => state.user.phone);
  const name = useAppSelector((state) => state.user.name);
  const email = useAppSelector((state) => state.user.email);

  console.log(phone, name, email);

  const handleSubmit = (e: SyntheticEvent) => {
    e.preventDefault();
  }

  return (
    <div className="flex flex-col gap-5 mb-25">
      <h3 className="text-3xl font-bold">Личные данные</h3>
      <form onSubmit={handleSubmit} className="flex flex-col gap-5">
        <DataInput label="Имя" value={name} canChange onSave={changeName} />
        <DataInput label="Номер телефона" value={formatPhoneFromE164(phone)} canChange={false} />
        <DataInput label="Почта" value={email} canChange onSave={changeEmail} />
      </form>
    </div>
  );
};

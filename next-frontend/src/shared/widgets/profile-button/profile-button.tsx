import { User } from "lucide-react";
import { Button } from "../../ui/button";

export const ProfileButton = () => {
  const isAuth = false; // TODO: Взять из стора

  return (
    <Button variant="outline" className="p-5">
      <User width={12} />
      <span>
        { isAuth ? "Профиль" : "Войти" }
      </span>
    </Button>
  )
}
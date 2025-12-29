import { useAppDispatch } from "@/app/store";
import { fetchLogout } from "../../store/userSlice";
import { useRouter } from "next/navigation";
import { clearCart } from "@/entities/cart/store/cartSlice";

export const LogoutButton = () => {
  const router = useRouter();
  const dispatch = useAppDispatch();

  const handleLogout = async () => {
    await dispatch(fetchLogout());

    dispatch(clearCart());
    
    router.push("/");
    router.refresh();
  };

  return (
    <button
      className="cursor-pointer mb-15 bg-gray-100 hover:bg-gray-300 transition-colors duration-150 linear px-5 py-2 rounded-2xl text-black"
      onClick={handleLogout}
    >
      Выйти
    </button>
  );
};

import { CartButton } from "@/shared/widgets/cart-button";
import { ProfileButton } from "@/shared/widgets/profile-button";
import { SearchInput } from "../search-input";

export const Header = () => {
  return (
    <>
      <header className="flex gap-5 py-10">
        <a href="/" className="flex gap-5 items-center">
          <img src="/logo.png" alt="logo" className="h-[35px]" />
          <div className="flex-col gap-3">
            <h1 className="uppercase font-[900] text-2xl leading-[100%]">
              Next pizza
            </h1>
            <p className="opacity-[0.5] text-[16px]">вкусней уже некуда</p>
          </div>
        </a>
        <SearchInput />

        <div className="flex items-center gap-2">
          <ProfileButton />
          <CartButton />
        </div>
      </header>
      <hr className="absolute left-0 right-0 h-[1px] bg-[#E5E5E5]" />
    </>
  );
};

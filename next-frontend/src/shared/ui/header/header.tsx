import { ProfileButton } from "@/shared/widgets/profile-button";
import { SearchInput } from "../search-input";
import { Logo } from "../logo";
import { Switch } from "../switch/switch";
import { ThemeSwitch } from "../switch/theme-switch";

export const Header = () => {
  return (
    <>
      <header className="flex gap-5 py-10">
        <Logo />
        <SearchInput />

        <div className="flex items-center gap-2">
          <ProfileButton />
          {/* <div className="flex flex-col gap-1 items-center text-sm">
            <Switch
              onClick={() => {
                localStorage.setItem(
                  "theme",
                  localStorage.getItem("theme") === "white" ? "dark" : "white"
                );
              }}
            />
            Сменить тему
          </div> */}
          <ThemeSwitch />
        </div>
      </header>
      <hr className="absolute left-0 right-0 h-[1px] bg-[#E5E5E5]" />
    </>
  );
};

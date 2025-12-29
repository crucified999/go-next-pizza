export const Logo = () => {
  return (
    <a href="/" className="flex gap-3 items-center">
      <img src="/logo.png" alt="logo" className="h-[40px]" />
      <div className="flex-col gap-3">
        <h1 className="uppercase font-[900] text-2xl leading-[100%]">
          Next pizza
        </h1>
        <p className="opacity-[0.5] text-[16px]">вкусней уже некуда</p>
      </div>
    </a>
  );
};

import Image from "next/image";

type PizzaImageProps = {
  src: string;
  alt: string;
  scale: number;
};

export const PizzaImage = ({ src, alt, scale }: PizzaImageProps) => {
  return (
    <div className="flex justify-center items-center pl-10 w-full h-full">
      <div
        className="transition-transform duration-300 ease-out will-change-transform"
        style={{ transform: `scale(${scale})` }}
      >
        <Image
          className="object-contain object-center"
          src={src}
          alt={alt}
          width={480}
          height={480}
        />
      </div>
    </div>
  );
};

export const Footer = () => {
  return (
    <footer className="bg-black text-white absolute right-0 left-0 flex justify-between p-10">
      <div className="flex-col gap-5">
        <h5 className="text-grey opacity-50">Партнерам</h5>
        <ul>
          <li>
            <a href="https://dodofranchise.ru/" target="_blank" className="hover:text-grey">Франшиза</a>
          </li>
          <li>
            <a href="https://dodobrands.io/investors/" target="_blank">Инвестиции</a>
          </li>
          <li>
            <a href="https://bidzaar.com/requests/public/buy?sorting.key=publishDate&sorting.direction=desc&logic=and&filters%5B0%5D.operator=in&filters%5B0%5D.field=companyId&filters%5B0%5D.value=%5Bd1827205-1309-4442-8adb-f6887474e684%5D" target="_blank">Поставщикам</a>
          </li>
          <li>
            <a href="https://www.dodoarenda.ru/" target="_blank">Предложить помещение</a>
          </li>
        </ul>
      </div>
      <div>
        <h5 className="text-grey opacity-50">Это интересно</h5>
        <ul>
          <li>
            <a href="https://www.dodokids.ru/" target="_blank">Экскурскии и мастер-классы</a>
          </li>
          <li>
            <a href="https://bezperchatok.ru/" target="_blank">Почему мы готовим без перчаток?</a>
          </li>
        </ul>
      </div>
      <div>
        <h5 className="text-grey opacity-50">Контакты</h5>
        <ul>
          <li>
            <a href="tel:88003020060">8 800 302-00-60</a>
          </li>
          <li>
            <a href="mailto:khalinarseny@yandex.ru">khalinarseny@yandex.ru</a>
          </li>
        </ul>
      </div>
      <hr />
      <div>
        <p className="text-grey opacity-50">© 2025 Dodo Pizza</p>
        <p className="text-grey opacity-50">Ни одно из прав не защищено</p>
        <ul>
          <li>
            <a href="https://www.instagram.com/dodopizza/" target="_blank">Instagram</a>
          </li>
          <li>
            <a href="https://www.facebook.com/dodopizza/" target="_blank">Facebook</a>
          </li>
        </ul>
      </div>
    </footer>
  );
};

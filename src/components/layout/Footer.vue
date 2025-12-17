<script setup>
import yugraLogo from '@/assets/footer/footer-yugra-logo.svg'
</script>

<template>
  <div class="footer-wrap">
    <footer class="footer">
      <div class="footer__container">

        <!-- ===== DESKTOP (как было) ===== -->
        <div class="footer__inner footer__desktop">
          <div class="footer__col">
            <h4 class="footer__title footer__title--small">Меню</h4>
            <ul class="footer__list">
              <li>Новости</li>
              <li>Сведения об образовательном учреждении</li>
              <li>Дополнительная информация</li>
              <li>Фотогалерея</li>
              <li>Странички педагогов</li>
              <li>Контакты</li>
            </ul>
          </div>

          <div class="footer__col">
            <h4 class="footer__title">Свежие новости</h4>
            <ul class="footer__list">
              <li>Задача организации</li>
              <li>Значимость этих проблем</li>
              <li>Повседневная практика показывает</li>
              <li>Идейные соображения высшего порядка</li>
              <li>С другой стороны</li>
            </ul>
          </div>

          <div class="footer__col">
            <h4 class="footer__title">Связаться с нами</h4>
            <p class="footer__text">8-800-800-80-80</p>
            <p class="footer__text">8-800-800-80-80</p>
          </div>
        </div>

        <!-- ===== MOBILE (компактно) ===== -->
        <div class="footer__mobile">
          <div class="footer__mobile-top">
            <h4 class="footer__title">Связаться с нами</h4>

            <div class="footer__phones">
              <a class="footer__phone" href="tel:88008008080">8-800-800-80-80</a>
              <a class="footer__phone" href="tel:88008008080">8-800-800-80-80</a>
            </div>

            <div class="footer__actions">
              <a class="footer__pill" href="#menu">Меню</a>
              <a class="footer__pill" href="#news">Новости</a>
              <a class="footer__pill" href="#contacts">Контакты</a>
            </div>
          </div>
        </div>

        <!-- ===== BOTTOM (общий) ===== -->
        <div class="footer__bottom">
          <img class="footer__logo" :src="yugraLogo" alt="Югра" />
          <p class="footer__license">
            Лицензия Департамента образования и науки ХМАО-Югры<br />
            №Л035-01304-86/00640475 от 09.02.2023 г.
          </p>
        </div>

      </div>
    </footer>
  </div>
</template>

<style scoped>
.footer__mobile { display: none; }
.footer__desktop { display: flex; }

@media (max-width: 800px) {
  .footer__desktop { display: none; }
  .footer__mobile { display: block; }

  .footer__mobile-top {
    padding: 10px 0 18px;
  }

  .footer__phones {
    display: flex;
    flex-direction: column;
    gap: 8px;
    margin-top: 10px;
  }

  .footer__phone {
    font-family: "Montserrat", system-ui, sans-serif;
    font-weight: 600;
    font-size: 18px;
    color: #FFFFFF;
    text-decoration: none;
  }

  .footer__actions {
    margin-top: 14px;
    display: flex;
    flex-wrap: wrap;
    gap: 10px;
  }

  .footer__pill {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    padding: 10px 14px;
    border-radius: 12px;
    background: rgba(255, 255, 255, 0.12);
    color: #fff;
    text-decoration: none;

    font-family: "Montserrat", system-ui, sans-serif;
    font-weight: 600;
    font-size: 14px;
  }
}
</style>

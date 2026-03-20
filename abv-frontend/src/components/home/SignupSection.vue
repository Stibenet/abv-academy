<template>
  <section class="signup">
    <div class="signup__inner">
      <div class="signup__card">
        <div class="signup__text">
          <span class="signup__text-main">
            Хотите узнать и увидеть как проходят<br/>
            занятия в нашей академии?
          </span>

          <span class="signup__text-sub">
            Тогда записывайтесь на личную встречу
          </span>
        </div>


        <form class="signup__form" @submit.prevent>
          <input class="signup__input" type="text" placeholder="Ваше имя"/>
          <input class="signup__input" type="tel" placeholder="Номер телефона"/>
          <button class="signup__btn" type="submit">ОСТАВИТЬ ЗАЯВКУ</button>
        </form>
      </div>
    </div>
  </section>
</template>

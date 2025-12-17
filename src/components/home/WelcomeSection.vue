<script setup>
import subtitleBrush from '@/assets/welcome/welcome-subtitle-brush.svg'
import brushMobile from '@/assets/hero/brush-mobile.svg'
</script>

<template>
  <section class="welcome">
    <div class="welcome__inner">

      <!-- Заголовок -->
      <h2 class="welcome__title">
        ДОБРО ПОЖАЛОВАТЬ В
        <span class="welcome__title-accent">ШКОЛУ «АБВ»</span>
      </h2>

      <!-- Подзаголовок с мазком -->
      <div class="welcome__subtitle">
        <picture class="welcome__subtitle-brush">
          <source :srcset="brushMobile" media="(max-width: 800px)" />
          <img :src="subtitleBrush" alt="" aria-hidden="true" />
        </picture>

        <span class="welcome__subtitle-text">Каждый ребенок – главная история успеха</span>
      </div>


      <!-- Текст -->
      <p class="welcome__text">
        В семейной атмосфере Академии ваш ребенок с 1 по 11 класс получит не просто знания,
        а фундамент для жизни: глубокое понимание предметов, разностороннее развитие личности
        и ключевые навыки XXI века, а опытные и вдохновляющие педагоги помогут ему раскрыть
        свои большие возможности
      </p>

    </div>
  </section>
</template>

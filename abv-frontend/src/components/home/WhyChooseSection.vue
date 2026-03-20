<script setup>
import whyWave from '@/assets/why/why-wave.svg'

import flexible from '@/assets/why/why-flexible.svg'
import smallClasses from '@/assets/why/why-small-classes.svg'
import english from '@/assets/why/why-english.svg'
import creative from '@/assets/why/why-creative.svg'
import childChoice from '@/assets/why/why-child-choice.svg'
import homework from '@/assets/why/why-homework.svg'
import tutorBadge from '@/assets/why/why-tutor-badge.svg'
</script>

<template>
  <section class="why">
    <!-- волна ВНЕ сетки -->
    <img class="why__wave" :src="whyWave" alt="" aria-hidden="true" />

    <div class="why__inner">
      <h2 class="why__title">
        ПОЧЕМУ ВАМ СТОИТ ВЫБРАТЬ
        <span class="why__title-accent">нашу школу</span>
      </h2>

      <div class="why__grid">

        <!-- ROW 1 -->
        <div class="why__row">
          <article class="why__card why__card--left">
            <img class="why__icon why__icon--1" :src="flexible" alt="" />
            <p class="why__card-text">Гибкая система<br />обучения</p>
          </article>

          <article class="why__card why__card--mid">
            <img class="why__icon why__icon--2" :src="smallClasses" alt="" />
            <p class="why__card-text">Небольшие классы<br />до 9 человек</p>
          </article>

          <article class="why__card why__card--right">
            <img class="why__icon why__icon--3" :src="english" alt="" />
            <p class="why__card-text">Английский<br />с 1 класса</p>
          </article>
        </div>

        <!-- ROW 2 -->
        <div class="why__row why__row--gap">
          <article class="why__card why__card--left">
            <img class="why__icon why__icon--4" :src="creative" alt="" />
            <p class="why__card-text">
              Интересные,<br />
              развивающие<br />
              и творческие<br />
              занятия
            </p>
          </article>

          <article class="why__card why__card--mid">
            <img class="why__icon why__icon--5" :src="childChoice" alt="" />
            <p class="why__card-text">
              Каждый ребёнок<br />
              сам выбирает<br />
              дополнительные<br />
              дисциплины
            </p>
          </article>

          <article class="why__card why__card--right">
            <img class="why__icon why__icon--6" :src="homework" alt="" />
            <p class="why__card-text">
              Домашнее задание<br />
              ребёнок делает под<br />
              присмотром<br />
              педагога
            </p>
          </article>
        </div>

        <!-- BOTTOM -->
        <div class="why__bottom">
          <p class="why__bottom-text">
            В течение года наставник-тьютор сопровождает ребёнка<br />
            и помогает добиться академических успехов
          </p>
          <img
              class="why__bottom-icon why__bottom-icon--tutor"
              :src="tutorBadge"
              alt=""
              aria-hidden="true"
          />
        </div>

      </div>
    </div>
  </section>
</template>

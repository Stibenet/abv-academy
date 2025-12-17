<script setup>
import { ref, onMounted, onBeforeUnmount, nextTick } from "vue";

import Header from "@/components/layout/Header.vue";
import Footer from "@/components/layout/Footer.vue";

import HeroSection from "@/components/home/HeroSection.vue";
import WelcomeSection from "@/components/home/WelcomeSection.vue";
import WhyChooseSection from "@/components/home/WhyChooseSection.vue";
import SignupSection from "@/components/home/SignupSection.vue";

const BASE_W = 1920;

const stageRef = ref(null);
const footerStageRef = ref(null);

const scale = ref(1);
const stageH = ref(0);
const footerH = ref(0);

let ro;

function recalc() {
  const vw = document.documentElement.clientWidth

  // 🔥 MOBILE MODE
  if (vw < 768) {
    scale.value = 1
    stageH.value = stageRef.value?.scrollHeight || 0
    footerH.value = footerStageRef.value?.scrollHeight || 0
    return
  }

  // DESKTOP SCALE MODE
  let raw = vw >= BASE_W ? 1 : vw / BASE_W
  const dpr = window.devicePixelRatio || 1
  raw = Math.round(raw * dpr) / dpr
  scale.value = raw

  requestAnimationFrame(() => {
    stageH.value = Math.ceil(stageRef.value?.scrollHeight * scale.value || 0)
    footerH.value = Math.ceil(footerStageRef.value?.scrollHeight * scale.value || 0)
  })
}


onMounted(async () => {
  await nextTick();
  recalc();

  ro = new ResizeObserver(recalc);
  ro.observe(document.documentElement);

  stageRef.value && ro.observe(stageRef.value);
  footerStageRef.value && ro.observe(footerStageRef.value);

  window.addEventListener("resize", recalc);
  window.addEventListener("load", recalc);
  document.fonts?.ready?.then(recalc);
});

onBeforeUnmount(() => {
  ro?.disconnect();
  window.removeEventListener("resize", recalc);
  window.removeEventListener("load", recalc);
});
</script>

<template>
  <div class="app-viewport">
    <!-- spacer держит общую высоту страницы -->
    <div class="app-spacer" :style="{ height: `${stageH + footerH}px` }" />

    <!-- ===== ОСНОВНОЙ КОНТЕНТ ===== -->
    <div class="stage-wrap">
      <div
          ref="stageRef"
          class="app-stage"
          :style="{ transform: `scale(${scale})` }"
      >
        <Header />
        <HeroSection />
        <WelcomeSection />
        <WhyChooseSection />
        <SignupSection />
      </div>
    </div>

    <!-- ===== ФУТЕР ===== -->
    <footer
        class="footer-shell"
        :style="{ top: `${stageH}px`, height: `${footerH}px` }"
    >
      <div class="footer-wrap">
        <div
            ref="footerStageRef"
            class="footer-stage"
            :style="{ transform: `scale(${scale})` }"
        >
          <Footer />
        </div>
      </div>
    </footer>
  </div>
</template>

<style scoped>
.app-viewport {
  position: relative;
  width: 100%;
  overflow-x: hidden;
}

.app-spacer {
  width: 1px;
}

/* ===== ОСНОВНОЙ STAGE ===== */
.stage-wrap {
  position: absolute;
  left: 50%;
  top: 0;
  transform: translateX(-50%);
  width: 1920px;
  height: 0; /* ключ: высотой управляет spacer */
}

.app-stage {
  position: absolute;
  left: 0;
  top: 0;
  width: 1920px;
  transform-origin: top center;
  will-change: transform;
}

/* ===== FOOTER ===== */
.footer-shell {
  position: absolute;
  left: 0;
  width: 100%;
  background: #4c4c4c;
  overflow: hidden;
}

.footer-wrap {
  position: absolute;
  left: 50%;
  top: 0;
  transform: translateX(-50%);
  width: 1920px;
  height: 0; /* ключ: высотой управляет footerH */
}

.footer-stage {
  position: absolute;
  left: 0;
  top: 0;
  width: 1920px;
  transform-origin: top center;
  will-change: transform;
}

</style>

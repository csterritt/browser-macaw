<template>
  <div class="mb-4 text-xl font-bold italic">About this application</div>

  <div class="ml-4">
    <div class="mb-2 text-lg">
      This application is just a different "front end" (the bit you see) the
      wonderful:
    </div>

    <div class="flex flex-row items-baseline mb-2">
      <span class="mr-8 text-xl font-bold">BrowserParrot application</span>

      <button class="btn btn-sm btn-primary" @click="visitBrowserParrot">
        Visit!
      </button>
    </div>

    <div class="mb-2 text-lg">
      Notably, BrowserParrot builds the database that this program uses... I
      just got on my SQL shoes and put a more complicated (but powerful!) user
      interface on BrowserParrot's data.
    </div>

    <div class="mb-2 text-lg">
      <span class="font-bold">HELP TEXT COMING SOON</span> but I sincerely hope
      that the text above the "Search" tab's entry fields are comprehensible.
    </div>
  </div>

  <div class="divider"></div>

  <div class="mb-4 text-xl font-bold italic">
    Resources used by this program
  </div>

  <div class="ml-4">
    <div class="flex flex-row items-center mb-2 text-lg">
      <div class="mr-8">
        This program's icon is derived from an icon in the Google Noto Emoji
        font.
        <br />
        It is used under the
        <span class="italic">Apache License Version 2.0</span>.
      </div>

      <button class="btn btn-sm btn-primary" @click="visitApacheLicense">
        See the license
      </button>
    </div>

    <div class="flex flex-row items-baseline mb-2 text-lg">
      <div class="mr-8">
        See the website https://github.com/googlefonts/noto-emoji for the entire
        font.
      </div>

      <button class="btn btn-sm btn-primary" @click="visitNotoEmoji">
        Get the font
      </button>
    </div>
  </div>
</template>

<script setup>
import { BrowserOpenURL } from '../../wailsjs/runtime'

const visitBrowserParrot = () => {
  BrowserOpenURL('https://browserparrot.com/')
}

const visitNotoEmoji = () => {
  BrowserOpenURL('https://github.com/googlefonts/noto-emoji')
}

const visitApacheLicense = () => {
  BrowserOpenURL('https://www.apache.org/licenses/LICENSE-2.0')
}
</script>

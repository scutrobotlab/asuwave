<template>
  <div class="text-center">
    <v-menu
      v-model="menu"
      :close-on-content-click="false"
      transition="slide-x-transition"
      top
      offset-y
    >
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          class="mb-8"
          color="secondary"
          dark
          absolute
          bottom
          left
          fab
          v-on="on"
          v-bind="attrs"
        >
          <v-icon>mdi-application-variable</v-icon>
        </v-btn>
      </template>
      <v-card>
        <v-card-title>调参面板</v-card-title>
        <ErrorAlert v-model="error" />
        <v-list dense>
          <v-list-item-group color="primary">
            <v-list-item v-for="i in variables" :key="i.Name">
              <v-list-item-content>
                <v-text-field
                  dense
                  v-model="i.Data"
                  v-bind:label="i.Name"
                  v-on:keyup.enter="modiVariable(i)"
                ></v-text-field>
              </v-list-item-content>
              <v-list-item-icon>
                <v-btn icon v-on:click="modiVariable(i)">
                  <v-icon>mdi-send</v-icon>
                </v-btn>
              </v-list-item-icon>
            </v-list-item>
          </v-list-item-group>
        </v-list>
      </v-card>
    </v-menu>
  </div>
</template>

<script>
import errorMixin from "@/mixins/errorMixin.js";
import { putVariable } from "@/api/variable.js";
export default {
  mixins: [errorMixin],
  data: () => ({
    menu: false,
  }),
  computed: {
    variables() {
      return this.$store.state.variables["modi"];
    },
  },
  methods: {
    openMenu() {
      this.menu = true;
    },
    async modiVariable(i) {
      await this.errorHandler(putVariable("modi", 1, i.Name, i.Type, i.Addr, parseFloat(i.Data)));
    },
  },
};
</script>

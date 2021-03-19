<template>
  <v-dialog v-model="dialog" max-width="400">
    <v-card>
      <v-list-item>
        <v-list-item-avatar tile size="125">
          <img src="@/assets/logo.png" alt="Logo" />
        </v-list-item-avatar>
        <v-list-item-content>
          <v-list-item-title>华工机器人实验室巨献</v-list-item-title>
          <v-list-item-subtitle>SCUT Robot Lab Present</v-list-item-subtitle>
        </v-list-item-content>
        <v-list-item-icon>
          <v-btn icon v-on:click="dialog = false">
            <v-icon>mdi-close</v-icon>
          </v-btn>
        </v-list-item-icon>
      </v-list-item>
      <v-card-text></v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn icon href="https://github.com/scutrobotlab/asuwave">
          <v-icon>mdi-github</v-icon>
        </v-btn>
        <v-btn icon v-on:click="hc = 'pink'" v-bind:color="hc">
          <v-icon>mdi-heart</v-icon>
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  data: () => ({
    dialog: false,
    hc: "",
  }),
  methods: {
    openDialog() {
      this.dialog = true;
    },
  },
};
</script>

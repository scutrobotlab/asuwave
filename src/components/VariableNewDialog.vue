<template>
  <v-row justify="center">
    <v-dialog v-model="dialog" persistent max-width="600px">
      <v-card>
        <v-toolbar dense color="primary">
          <v-toolbar-title>添加变量</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-btn icon dark @click="dialog = false">
            <v-icon>mdi-close</v-icon>
          </v-btn>
        </v-toolbar>
        <v-card-text>
          <v-container>
            <ErrorAlert v-model="error" />
            <v-form ref="form" v-model="valid" lazy-validation>
              <v-row>
                <v-col cols="12" sm="6" md="6">
                  <v-select
                    :items="[1, 2, 3]"
                    label="板子代号"
                    hint="保持默认为1即可"
                    :rules="[(v) => !!v || '板子代号是必要的']"
                    required
                    v-model="Board"
                  ></v-select>
                </v-col>
                <v-col cols="12" sm="6" md="6">
                  <v-select
                    :items="types"
                    label="变量类型"
                    :rules="[(v) => !!v || '变量类型是必要的']"
                    required
                    v-model="Type"
                  ></v-select>
                </v-col>
                <v-col cols="12" sm="6" md="6">
                  <v-text-field
                    label="变量名"
                    type="text"
                    :rules="[(v) => !!v || '变量名是必要的']"
                    required
                    v-model="Name"
                  ></v-text-field>
                </v-col>
                <v-col cols="12" sm="6" md="6">
                  <v-text-field
                    label="变量地址"
                    type="text"
                    :rules="AddrRules"
                    hint="形如2000ab78"
                    required
                    v-model="Addr"
                  ></v-text-field>
                </v-col>
              </v-row>
            </v-form>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="error" @click="$refs.form.reset()">清空</v-btn>
          <v-btn color="primary" :disabled="!valid" @click="addVariable()">添加</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
import errorMixin from "@/mixins/errorMixin.js";
import { postVariable } from "@/api/variable.js";

export default {
  mixins: [errorMixin],
  props: ["opt"],
  data: () => ({
    dialog: false,
    valid: true,
    Board: 1,
    Name: "",
    Type: "",
    Addr: "",
    AddrRules: [
      (v) => !!v || "变量地址是必要的",
      (v) => /2[0-9a-f]{7}/.test(v) || "格式错误，应形如2000ab78",
    ],
  }),
  computed: {
    types() {
      return this.$store.state.vTypes;
    },
  },
  methods: {
    openDialog() {
      this.dialog = true;
    },
    addVariable() {
      if (this.$refs.form.validate()) {
        this.errorHandler(
          postVariable(this.opt, 1, this.Name, this.Type, parseInt(this.Addr, 16))
        ).then(async () => {
          this.dialog = false;
          await this.$store.dispatch("getV", this.opt);
        });
      }
    },
  },
};
</script>

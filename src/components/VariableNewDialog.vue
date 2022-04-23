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
                  <v-color-picker
                    dot-size="25"
                    swatches-max-height="100"
                    show-swatches
                    v-model="Inputcolor"
                  ></v-color-picker>
                </v-col>
                <v-col cols="12" sm="6" md="6">
                  <v-select
                    :items="[1, 2, 3]"
                    label="板子代号"
                    hint="保持默认为1即可"
                    :rules="[(v) => !!v || '板子代号是必要的']"
                    required
                    v-model="Board"
                  ></v-select>
                  <v-row>
                    <v-col cols="12" sm="12" md="12">
                      <v-select
                        :items="types"
                        label="变量类型"
                        :rules="[(v) => !!v || '变量类型是必要的']"
                        required
                        :disabled="disable"
                        v-model="Type"
                      ></v-select>
                    </v-col>
                  </v-row>

                  <v-row>
                    <v-col cols="12" sm="12" md="12">
                      <v-text-field
                        label="变量名称"
                        type="text"
                        :rules="[(v) => !!v || '变量名是必要的']"
                        required
                        v-model="Name"
                      ></v-text-field>
                    </v-col>
                  </v-row>
                  <v-row>
                    <v-col cols="12" sm="12" md="12">
                      <v-text-field
                        label="变量地址"
                        type="text"
                        :rules="AddrRules"
                        :disabled="disable"
                        hint="[2000 0000, 2fff ffff]区间的16进制数。"
                        required
                        v-model="Addr"
                      ></v-text-field>
                    </v-col>
                  </v-row>
                  <v-row>
                    <v-col cols="12" sm="12" md="12">
                      <v-text-field
                        label="变量颜色"
                        type="text"
                        disabled
                        required
                        v-model="Inputcolor"
                      ></v-text-field>
                    </v-col>
                  </v-row>
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
    disable: false,
    colorcard: false,
    valid: true,
    Board: 1,
    Name: "",
    Type: "",
    Addr: "",
    Inputcolor: "#F3CB09FF",
    showcolorcard: "",
    AddrRules: [
      (v) => !!v || "变量地址是必要的",
      (v) => /^(0x)?2[0-9a-fA-F]{7}$/.test(v) || "请输入[2000 0000, 2fff ffff]区间的16进制数。",
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
    openDialogFromList(name, type, board, addr) {
      this.dialog = true;
      this.disable = true;
      this.Name = name;
      this.Type = type;
      this.Addr = addr;
    },
    addVariable() {
      if (this.$refs.form.validate()) {
        this.errorHandler(
          postVariable(this.opt, 1, this.Name, this.Type, parseInt(this.Addr, 16), this.Inputcolor)
        ).then(async () => {
          this.dialog = false;
          await this.$store.dispatch("getV", this.opt);
          this.$bus.$emit("sendcolor", this.Inputcolor);
          this.$bus.$emit("sendalert", true);
          console.log(this.opt);
        });
      }
    },
  },
};
</script>

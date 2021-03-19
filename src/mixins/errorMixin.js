import ErrorAlert from "@/components/ErrorAlert.vue";

export default {
  data: () => ({
    error: null,
  }),
  methods: {
    clearError() {
      this.error = null;
    },
    async errorHandler(func) {
      this.clearError();
      try {
        const res = await func;
        return res;
      } catch (error) {
        this.error = error;
        throw error;
      }
    },
  },
  components: {
    ErrorAlert,
  },
};

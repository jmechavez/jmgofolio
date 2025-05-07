document.addEventListener('alpine:init', () => {


  // Store + component for search box
  Alpine.store('search', {
    showSearch: () => {}
  });

  Alpine.data('searchComponent', () => ({
    visible: false,
    init() {
      Alpine.store('search').showSearch = this.showSearch.bind(this);
    },
    showSearch() {
      this.visible = true;
      this.$nextTick(() => {
        this.$refs.input.value = '';
        this.$refs.input.focus();
      });
    },
    hideSearch() {
      setTimeout(() => {
        this.visible = false;
      }, 100);
    },
    blurInput() {
      this.$refs.input.blur();
    }
  }));
});

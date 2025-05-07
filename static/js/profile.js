// Alpine.js component
console.log("Profile.js loaded");
function profileData() {
  console.log("profileData function called");
  return {
    profile: null,
    loading: true,
    error: null,
    fetchProfile() {
      console.log("fetchProfile called");
      this.loading = true;
      fetch('/api/portfolio')
        .then(res => {
          console.log("Fetch response received", res);
          if (!res.ok) throw new Error('Failed to load profile data');
          return res.json();
        })
        .then(data => {
          console.log("Profile data loaded", data);
          this.profile = data;
          this.loading = false;
        })
        .catch(err => {
          console.error("Error in fetch:", err);
          this.error = err.message;
          this.loading = false;
        });
    }
  }
}

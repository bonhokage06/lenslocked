class Nav {
  constructor() {
    this.mobileNavClose = document.getElementById("mobile-nav-close")
    this.mobileNav = document.getElementById("mobile-nav")
    this.mobileHamburger = document.getElementById("mobile-hamburger")
    this.mobileHamburger.addEventListener("click", () => {
      this.toggle()
    })
    this.mobileNavClose.addEventListener("click", () => {
      this.toggle()
    })
  }
  toggle() {
    this.mobileNav.classList.toggle("invisible")
    this.mobileNav.classList.toggle("visible")
    if (this.mobileNav.classList.contains("visible")) {
      this.mobileNav.classList.toggle("animate-fade")
    }
  }
}

export default Nav

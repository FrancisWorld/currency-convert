# 💱 Go Currency Converter

Welcome to the Go Currency Converter, a simple yet powerful command-line tool for converting currencies using real-time exchange rates!

## 🌟 Features

- Convert between 5 major currencies: USD, EUR, GBP, JPY, and BRL
- Real-time exchange rates from reliable API sources
- User-friendly command-line interface
- Error handling and input validation
- Fallback API for increased reliability

## 🚀 Getting Started

### Prerequisites

- Go 1.16 or higher
- Internet connection for fetching exchange rates

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/go-currency-converter.git
   ```

2. Navigate to the project directory:
   ```
   cd go-currency-converter
   ```

3. Install dependencies:
   ```
   go mod tidy
   ```

### Usage

Run the program:
```
go run main.go
```


Follow the on-screen prompts to:
1. Select the source currency
2. Select the target currency
3. Enter the amount to convert
4. View the conversion result
5. Choose to perform another conversion or exit

## 🛠️ Technologies Used

- [Go](https://golang.org/) - The programming language used
- [huh](https://github.com/charmbracelet/huh) - For creating interactive command-line interfaces
- [Currency API](https://github.com/fawazahmed0/currency-api) - Primary source for exchange rates
- [Currency API (Backup)](https://currency-api.pages.dev/) - Secondary source for exchange rates

## 📚 What I Learned

This project helped me gain hands-on experience with:

- Go programming fundamentals
- Working with external APIs in Go
- Error handling and input validation
- Creating user-friendly CLI applications
- Implementing fallback mechanisms for increased reliability

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](https://github.com/yourusername/go-currency-converter/issues).

## 📝 License

This project is [MIT](https://choosealicense.com/licenses/mit/) licensed.

## 🙏 Acknowledgements

- [Currency API](https://github.com/fawazahmed0/currency-api) for providing free and reliable exchange rate data
- [Charm](https://charm.sh/) for the amazing `huh` library

---

Happy converting! 💰✨

package q1


import "fmt"


func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
  var soma float64
  var discount float64
  for _, num := range purchaseHistory {
     soma += num


  }
  if soma == 0 {
     discount = currentPurchase * 0.1
  } else if soma/float64(len(purchaseHistory)) > 1000 {
     discount = currentPurchase * 0.2
  } else if soma > 1000 {
     discount = currentPurchase * 0.1
  } else if soma <= 1000 && soma > 500 {
     discount = currentPurchase * 0.05
  } else if soma <= 500 {
     discount = currentPurchase * 0.02
  }



  return 0, nil
}

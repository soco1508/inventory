package handler

import (
	"backend/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	productService           service.ProductService
	saleSummaryService       service.SaleSummaryService
	purchaseSummaryService   service.PurchaseSummaryService
	expenseSummaryService    service.ExpenseSummaryService
	expenseByCategoryService service.ExpenseByCategoryService
}

func NewDashboardHandler(
	productService service.ProductService,
	saleSummaryService service.SaleSummaryService,
	purchaseSummaryService service.PurchaseSummaryService,
	expenseSummaryService service.ExpenseSummaryService,
	expenseByCategoryService service.ExpenseByCategoryService,
) *DashboardHandler {
	return &DashboardHandler{
		productService:           productService,
		saleSummaryService:       saleSummaryService,
		purchaseSummaryService:   purchaseSummaryService,
		expenseSummaryService:    expenseSummaryService,
		expenseByCategoryService: expenseByCategoryService,
	}
}

func (h *DashboardHandler) GetDashboardMetrics(c *gin.Context) {
	ctx := c.Request.Context()
	var dashboardMetrics []any
	products, err := h.productService.GetPopularProducts(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Fatalf("%v", err)
	}
	dashboardMetrics = append(dashboardMetrics, products)

	saleSummary, err := h.saleSummaryService.GetSaleSummary(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Fatalf("%v", err)
	}
	dashboardMetrics = append(dashboardMetrics, saleSummary)

	purchaseSummary, err := h.purchaseSummaryService.GetPurchaseSummary(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Fatalf("%v", err)
	}
	dashboardMetrics = append(dashboardMetrics, purchaseSummary)

	expenseSummary, err := h.expenseSummaryService.GetExpenseSummary(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Fatalf("%v", err)
	}
	dashboardMetrics = append(dashboardMetrics, expenseSummary)

	expenseByCategory, err := h.expenseByCategoryService.GetExpenseByCategory(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Fatalf("%v", err)
	}
	dashboardMetrics = append(dashboardMetrics, expenseByCategory)

	c.JSON(http.StatusOK, dashboardMetrics)
}

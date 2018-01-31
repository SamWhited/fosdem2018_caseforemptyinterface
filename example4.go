func AddMetrics(
	ctx context.Context,
	metrics prometheus.Registry,
) context.Context {
	/* … */
}

func AddDatabase(
	ctx context.Context,
	db *sql.DB,
) context.Context {
	/* … */
}

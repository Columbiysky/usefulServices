using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace DbCreatorAndMigrator.Migrations
{
    /// <inheritdoc />
    public partial class PairTablePriceColumnMadeDouble : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.AlterColumn<double>(
                name: "Price",
                schema: "finance",
                table: "Pair",
                type: "double precision",
                nullable: false,
                oldClrType: typeof(long),
                oldType: "bigint");
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.AlterColumn<long>(
                name: "Price",
                schema: "finance",
                table: "Pair",
                type: "bigint",
                nullable: false,
                oldClrType: typeof(double),
                oldType: "double precision");
        }
    }
}

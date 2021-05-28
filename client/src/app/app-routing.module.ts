import { HttpClientModule } from "@angular/common/http";
import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { RouterModule, Routes } from "@angular/router";
import { HomeComponent } from "./home/home.component";
import { LoginComponent } from "./login/login.component";
import { NotFoundComponent } from "./not-found/not-found.component";
import { OrbitalServerService } from "./orbital/orbital-server.service";

const routes: Routes = [
	{ path: "home", component: HomeComponent },
	{ path: "", component: LoginComponent },
	{ path: "**", component: NotFoundComponent },
];

@NgModule({
	imports: [BrowserModule, HttpClientModule, RouterModule.forRoot(routes)],
	providers: [OrbitalServerService],
	exports: [RouterModule],
})
export class AppRoutingModule {}

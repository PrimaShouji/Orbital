import { HttpClientModule } from "@angular/common/http";
import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { RouterModule, Routes } from "@angular/router";
import { LoginComponent } from "./components/login/login.component";
import { OrbitalServerService } from "./services/orbital/orbital-server.service";

const routes: Routes = [{ path: "", component: LoginComponent }];

@NgModule({
	imports: [BrowserModule, HttpClientModule, RouterModule.forRoot(routes)],
	providers: [OrbitalServerService],
	exports: [RouterModule],
})
export class AppRoutingModule {}
